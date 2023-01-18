package server

import (
	"context"
	"fmt"
	"net"

	"github.com/LTitan/Mebius/pkg/apis/v1alpha1"
	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

type ServerInterface interface {
	RegisterCommand()
}

type RawServer struct {
	opts *options.GlobalOption
}

func NewServer(opts *options.GlobalOption) ServerInterface {
	return &RawServer{
		opts: opts,
	}
}

func (rs *RawServer) GetMachine(ctx context.Context, req *v1alpha1.Machine) (resp *v1alpha1.Machine, err error) {
	return req, nil
}

func (rs *RawServer) RegisterCommand() {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "mebius grpc server",
		Long:  "mebius grpc server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rs.Run()
		},
	}
	rs.opts.GetCommand().AddCommand(cmd)
}

func (rs *RawServer) Run() error {
	// TODO: Add some grpc middleware
	ctx, cancel := mcontext.WithWaitGroup(context.Background()).WithCancel()
	defer cancel()
	server := grpc.NewServer(grpc.MaxMsgSize(1024 * 1024 * 5))
	protos.RegisterServerServer(server, rs)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", rs.opts.Sever().Port))
	if err != nil {
		return err
	}
	klog.Infof("start grpc server, listen on *:%d", rs.opts.Sever().Port)
	if err := server.Serve(listen); err != nil {
		return err
	}
	<-ctx.Done()
	server.Stop()
	return nil
}
