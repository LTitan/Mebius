package server

import (
	"math"

	"github.com/spf13/cobra"
)

const (
	connectionMaxSendAndRecvSize = math.MaxUint32
)

type Option struct {
	GlobalOption

	Port   int32
	Thread int32
}

type GlobalOption struct {
	EnableKubeConfig bool
	KubeConfig       string
}

func (g *GlobalOption) Parse(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&g.EnableKubeConfig, "enable-kubeconfig", true, "is enable kube-config")
	cmd.PersistentFlags().StringVar(&g.KubeConfig, "kubeconfig", "", "kube-config file")
}

func (o *Option) Parse(cmd *cobra.Command) {
	o.GlobalOption.Parse(cmd)
	cmd.Flags().Int32Var(&o.Port, "port", 8000, "listen port")
	cmd.Flags().Int32Var(&o.Thread, "thread", 8000, "thread size")
}
