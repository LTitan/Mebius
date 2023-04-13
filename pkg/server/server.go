package server

import (
	"context"
	"fmt"
	"net"

	"github.com/LTitan/Mebius/apis/v1alpha1"
	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/factory"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/protos"
	"github.com/LTitan/Mebius/pkg/protos/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type RawServer struct {
	factory.FrameworkInterface
	opts *options.GlobalOption
}

func NewServer(opts *options.GlobalOption) factory.Application {
	return &RawServer{
		opts:               opts,
		FrameworkInterface: factory.NewBaseFramework(opts),
	}
}

func (rs *RawServer) GetMachine(ctx context.Context, req *types.ExampleRequest) (resp *types.ExampleResponse, err error) {
	resp = &types.ExampleResponse{
		Content: "this is a grpc server",
	}
	return
}

func (rs *RawServer) Heartbeat(ctx context.Context, req *types.HeartbeatRequest) (resp *types.HeartbeatResponse, err error) {
	klog.Infof("%+v", *req)
	resp = &types.HeartbeatResponse{
		Code: 0,
	}
	machine, err := rs.GetMebiusClientSet().MebiusV1alpha1().Machines("default").Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		resp.Code = 4
		resp.Message = err.Error()
		err = nil
		return
	}
	resp.Data = []*v1alpha1.Machine{
		machine,
	}
	return
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
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(rs.opts.MaxRecvByteSize),
		grpc.MaxSendMsgSize(rs.opts.MaxSendByteSize),
	)
	if err := rs.Init(); err != nil {
		return err
	}
	klog.Infof("start grpc server, listen on *:%d", rs.opts.Sever().Port)
	protos.RegisterServerServer(server, rs)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", rs.opts.Sever().Port))
	if err != nil {
		return err
	}
	if err := server.Serve(listen); err != nil {
		return err
	}
	<-ctx.Done()
	server.Stop()
	return nil
}
