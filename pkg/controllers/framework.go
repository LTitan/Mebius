package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	mebiusclientset "github.com/LTitan/Mebius/pkg/clients/clientset/mebius"
	"github.com/LTitan/Mebius/pkg/clients/informer/externalversions"
	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Interface .
type Interface interface {
	Run(ctx mcontext.MContext, threadSize int) error
}

type Framework struct {
	kubeConfig                  *rest.Config
	mebiusSharedInformerFactory externalversions.SharedInformerFactory
	mebiusClientSet             mebiusclientset.Interface

	opts          *options.GlobalOption
	controllerSet map[string]Interface
}

func NewFramework(opts *options.GlobalOption) *Framework {
	f := &Framework{
		opts: opts,
	}
	return f
}

func (f *Framework) RegisterCommand() {
	cmd := &cobra.Command{
		Use:   "controller",
		Short: "mebius controller executor",
		Long:  "mebius controller executor",
		RunE: func(cmd *cobra.Command, args []string) error {
			return f.Run()
		},
	}
	f.opts.GetCommand().AddCommand(cmd)
}

func (f *Framework) buildKubeConfig() (err error) {
	f.kubeConfig, err = clientcmd.BuildConfigFromFlags("", f.opts.KubeConfig)
	return
}

func (f *Framework) buildClientSet() (err error) {
	f.mebiusClientSet, err = mebiusclientset.NewForConfig(f.kubeConfig)
	return
}

func (f *Framework) buildInformer() (err error) {
	f.mebiusSharedInformerFactory = externalversions.NewSharedInformerFactory(f.mebiusClientSet,
		time.Hour*time.Duration(f.opts.Controller().ResyncPeriod))
	return
}

func (f *Framework) buildControllerFactory() (err error) {
	// machine controller
	f.controllerSet = map[string]Interface{
		machineControllerName: NewMachineController(f.mebiusSharedInformerFactory.Mebius().
			V1alpha1().Machines()),
	}
	return nil
}

func (f *Framework) init() error {
	type initFunc func() error
	funcs := []initFunc{
		f.buildKubeConfig,
		f.buildClientSet,
		f.buildInformer,
		f.buildControllerFactory,
	}
	for _, realRun := range funcs {
		if err := realRun(); err != nil {
			return err
		}
	}
	return nil
}

func (f *Framework) getRegisteredControllerName() []string {
	res := []string{}
	for name := range f.controllerSet {
		res = append(res, name)
	}
	return res
}

func (f *Framework) Run() (err error) {
	if err := f.init(); err != nil {
		return err
	}
	if f.opts.Controller().Name == "" || f.controllerSet[f.opts.Controller().Name] == nil {
		return fmt.Errorf("not found %s controller, there are %s controllers", f.opts.Controller().Name,
			strings.Join(f.getRegisteredControllerName(), "|"))
	}
	stopCh := options.SetupSignalHandler()
	ctx, cancel := mcontext.WithWaitGroup(context.Background()).WithCancel()
	ctx.GO(func() { <-stopCh; cancel() })
	ctx.GO(func() {
		f.mebiusSharedInformerFactory.Start(ctx.Done())
	})
	return f.controllerSet[f.opts.Controller().Name].Run(ctx, f.opts.ThreadSize)
}
