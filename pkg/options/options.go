package options

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

const (
	defaultMaxRecvAndSendByteSize = 1024 * 1024 * 1024 * 2
)

type GlobalOption struct {
	EnableKubeConfig bool
	KubeConfig       string
	ThreadSize       int
	MaxRecvByteSize  int
	MaxSendByteSize  int

	cmd *cobra.Command

	copt    controllerOptions
	server  serverOptions
	gateway gatewayOptions
}

func (g *GlobalOption) Parse() {
	g.cmd.PersistentFlags().BoolVar(&g.EnableKubeConfig, "enable-kubeconfig", true, "is enable kube-config")
	g.cmd.PersistentFlags().StringVar(&g.KubeConfig, "kubeconfig", clientcmd.RecommendedHomeFile, "kube-config file")
	g.cmd.PersistentFlags().IntVar(&g.ThreadSize, "thread-size", 10, "size of controller thread")
	g.cmd.PersistentFlags().IntVar(&g.MaxSendByteSize, "max-send-byte-size", defaultMaxRecvAndSendByteSize,
		"max byte size of connection send by grpc/gateway")
	g.cmd.PersistentFlags().IntVar(&g.MaxRecvByteSize, "max-recv-byte-size", defaultMaxRecvAndSendByteSize,
		"max byte size of connection recv by grpc/gateway")

	g.cmd.PersistentFlags().StringVar(&g.copt.Name, "controller-name", "unknown", "start controller name")
	g.cmd.PersistentFlags().IntVar(&g.copt.ResyncPeriod, "resync", 24, "informer resync period hour")
	g.cmd.PersistentFlags().IntVar(&g.server.Port, "port", 8089, "grpc listen port")
	g.cmd.PersistentFlags().IntVar(&g.gateway.Port, "gateway-port", 8080, "gateway listen port")
	g.cmd.PersistentFlags().StringVar(&g.gateway.Endpoints, "grpc-endpoints", "localhost:8089", "grpc backend endpoints")
}

func (g *GlobalOption) ExecuteOrDie() {
	defer klog.Flush()
	// validate && running
	if err := g.Validate(); err != nil {
		klog.Fatalln(err)
	}

	if err := g.cmd.Execute(); err != nil {
		klog.Fatalln(err)
	}
}

func (g *GlobalOption) Validate() error {
	if g.EnableKubeConfig && g.KubeConfig == "" {
		return fmt.Errorf("kubeconfig path must be set when enbale kubeconfig")
	}
	if err := g.copt.Validate(); err != nil {
		return nil
	}
	return nil
}

func (g *GlobalOption) GetCommand() *cobra.Command {
	return g.cmd
}

func (g *GlobalOption) RegisteredRunE(run func(opt *GlobalOption) func(cmd *cobra.Command, args []string) error) {
	g.cmd.RunE = run(g)
}

func NewRootCommand() *GlobalOption {
	g := &GlobalOption{
		cmd: &cobra.Command{
			Use:   "mebius",
			Short: "mebius project executor",
			Long:  "mebius is a Kubernetes Operator practices, using mebius [COMMAND] [FLAGS] to start it",
			// TODO: We may need a subcommand to execute it
		},
	}
	klog.InitFlags(nil)
	g.Parse()
	return g
}

func (g *GlobalOption) Controller() *controllerOptions {
	return &g.copt
}

func (g *GlobalOption) Sever() *serverOptions {
	return &g.server
}

func (g *GlobalOption) Gateway() *gatewayOptions {
	return &g.gateway
}

// controller options
type controllerOptions struct {
	Name         string
	ResyncPeriod int
}

func (c *controllerOptions) Validate() error {
	return nil
}

type serverOptions struct {
	Port int
}

func (c *serverOptions) Validate() error {
	return nil
}

type gatewayOptions struct {
	Endpoints string
	Port      int
}

func (c *gatewayOptions) Validate() error {
	return nil
}
