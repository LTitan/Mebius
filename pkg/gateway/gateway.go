package gateway

import (
	"context"
	"fmt"
	"net/http"

	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/factory"
	apiv1 "github.com/LTitan/Mebius/pkg/gateway/api/v1"
	"github.com/LTitan/Mebius/pkg/gateway/middleware"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/protos"
	"github.com/LTitan/Mebius/pkg/utils/function"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"k8s.io/klog/v2"
)

type gateway struct {
	opts *options.GlobalOption
}

func NewGateway(opts *options.GlobalOption) factory.Application {
	return &gateway{
		opts: opts,
	}
}

func (rs *gateway) RegisterCommand() {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "mebius gateway server",
		Long:  "mebius gateway server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rs.Run()
		},
	}
	rs.opts.GetCommand().AddCommand(cmd)
}

func (g *gateway) Run() error {
	gin.SetMode(gin.ReleaseMode)
	ctx, cancel := mcontext.WithWaitGroup(context.Background()).WithCancel()
	defer cancel()
	mux := g.newRpcGatewayServerMux()
	err := protos.RegisterServerHandlerFromEndpoint(ctx, mux, g.opts.Gateway().Endpoints,
		g.rpcDialOptions())
	if err != nil {
		klog.Errorf("Register ServerHandler from endpoint %s error", g.opts.Gateway().Endpoints)
		return err
	}
	server := gin.New()
	g.router(server, mux)
	addr := fmt.Sprintf(":%d", g.opts.Gateway().Port)
	return g.registerDocAndStartServer(server, addr)
}

func (g *gateway) router(server *gin.Engine, mux *runtime.ServeMux) {
	server.Use(middleware.CORSMiddleware(), middleware.Logger()) //add more
	// warp grpc runtime mux to gin handlers
	server.Group("/api/*{v1}").Any("", gin.WrapH(mux))

	// apiv1 group:
	// Note /api/v1 is conflict with /api/v1alpha1, they have same prefix `api/v1`
	// handle more
	apiv1.RegisterRouter(server.Group("/apiv1"))

}

func (g *gateway) registerDocAndStartServer(server *gin.Engine, addr string) error {
	klog.Infof("start grpc server, listen on %s", addr)
	return function.NewFunctionLinkErr(
		func() error {
			return middleware.SwaggerDoc(server)
		},
		func() error {
			return server.Run(addr)
		}).DoErr()
}

func (g *gateway) rpcDialOptions() []grpc.DialOption {
	opts := []grpc.DialOption{
		// local credentials
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return opts
}

func (g *gateway) newRpcGatewayServerMux() *runtime.ServeMux {
	return runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			md := metadata.Pairs("auth", header)
			return md
		}),
	)
}
