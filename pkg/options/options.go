package options

import "github.com/spf13/cobra"

type GlobalOption struct {
	EnableKubeConfig bool
	KubeConfig       string

	Cmd *cobra.Command
}

func (g *GlobalOption) Parse() {
	g.Cmd.PersistentFlags().BoolVar(&g.EnableKubeConfig, "enable-kubeconfig", true, "is enable kube-config")
	g.Cmd.PersistentFlags().StringVar(&g.KubeConfig, "kubeconfig", "", "kube-config file")
}

func (g *GlobalOption) ExecuteOrDie() {
	if err := g.Cmd.Execute(); err != nil {
		panic(err)
	}
}

func (g *GlobalOption) Validate() error {
	return nil
}

func NewRootCommand() *GlobalOption {
	g := &GlobalOption{
		Cmd: &cobra.Command{
			Use:   "mebius",
			Short: "mebius project executor",
			Long:  "mebius is a Kubernetes Operator practices, using mebius [COMMAND] [FLAGS] to start it",
		},
	}
	g.Parse()
	return g
}
