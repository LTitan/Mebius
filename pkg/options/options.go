package options

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

type GlobalOption struct {
	EnableKubeConfig bool
	KubeConfig       string
	ThreadSize       int

	cmd *cobra.Command
}

func (g *GlobalOption) Parse() {
	g.cmd.PersistentFlags().BoolVar(&g.EnableKubeConfig, "enable-kubeconfig", true, "is enable kube-config")
	g.cmd.PersistentFlags().StringVar(&g.KubeConfig, "kubeconfig", clientcmd.RecommendedHomeFile, "kube-config file")
	g.cmd.PersistentFlags().IntVar(&g.ThreadSize, "thread-size", 10, "size of controller thread")
}

func (g *GlobalOption) ExecuteOrDie() {
	if err := g.cmd.Execute(); err != nil {
		panic(err)
	}
}

func (g *GlobalOption) Validate() error {
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
