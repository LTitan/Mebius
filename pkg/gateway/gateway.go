package gateway

import (
	"context"
	"fmt"
	"net/http"

	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/factory"
	"github.com/LTitan/Mebius/pkg/gateway/middleware"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/protos"
	"github.com/LTitan/Mebius/pkg/utils/function"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			md := metadata.Pairs("auth", header)
			return md
		}),
	)
	dialOption := []grpc.DialOption{grpc.WithInsecure()}
	protos.RegisterServerHandlerFromEndpoint(ctx, mux, g.opts.Gateway().Endpoints, dialOption)
	addr := fmt.Sprintf(":%d", g.opts.Gateway().Port)

	server := gin.New()
	server.Use(middleware.CORSMiddleware(), middleware.Logger()) //add more
	server.Group("/api/*{v1}").Any("", gin.WrapH(mux))
	server.GET("/apiv1/test", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	klog.Infof("start grpc server, listen on %s", addr)
	return function.NewFunctionLinkErr(
		func() error {
			return middleware.SwaggerDoc(server)
		},
		func() error {
			return server.Run(addr)
		}).DoErr()
}
