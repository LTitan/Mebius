package controllers

import (
	"context"
	"fmt"
	"strings"

	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/factory"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/utils/function"
	"github.com/spf13/cobra"
)

// Interface .
type Interface interface {
	Run(ctx mcontext.MContext, threadSize int) error
}

type Framework struct {
	factory.FrameworkInterface
	opts          *options.GlobalOption
	controllerSet map[string]Interface
}

func NewFramework(opts *options.GlobalOption) factory.Application {
	f := &Framework{
		opts:               opts,
		FrameworkInterface: factory.NewBaseFramework(opts),
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

func (f *Framework) buildControllerFactory() (err error) {
	// machine controller
	// 新增controller在这里注册
	f.controllerSet = map[string]Interface{
		machineControllerName: NewMachineController(f.GetMebiusSharedInformerFactory().Mebius().
			V1alpha1().Machines()),
	}
	return nil
}

func (f *Framework) init() error {
	fle := function.NewFunctionLinkErr(
		f.FrameworkInterface.Init,
		f.buildControllerFactory,
	)
	return fle.DoErr()
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
	ctx.Go(func() { <-stopCh; cancel() })
	ctx.Go(func() {
		f.GetMebiusSharedInformerFactory().Start(ctx.Done())
	})
	return f.controllerSet[f.opts.Controller().Name].Run(ctx, f.opts.ThreadSize)
}
