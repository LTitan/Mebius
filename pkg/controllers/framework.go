package controllers

import (
	"context"
	"time"

	mebiusclientset "github.com/LTitan/Mebius/pkg/clients/clientset/mebius"
	"github.com/LTitan/Mebius/pkg/clients/informer/externalversions"
	mcontext "github.com/LTitan/Mebius/pkg/context"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

// Interface .
type Interface interface {
	Run(ctx mcontext.MContext, threadSize int)
}

func RegisterFramework(opts *options.GlobalOption) {
	opts.RegisteredRunE(runController)
}

func runController(opt *options.GlobalOption) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// TODO: Deal with enable-kubeconfig
		kubeconfig := opt.KubeConfig
		// set up signals so we handle the first shutdown signal gracefully
		stopCh := options.SetupSignalHandler()
		ctx, cancel := mcontext.WithWaitGroup(context.Background()).WithCancel()
		ctx.GO(func() { <-stopCh; cancel() })
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			klog.Fatalln(err)
			return err
		}

		clientset, err := mebiusclientset.NewForConfig(config)
		if err != nil {
			klog.Fatalln(err)
			return err
		}

		// factory init .
		factory := externalversions.NewSharedInformerFactory(clientset, time.Second*30)
		machineController := NewMachineController(factory.Mebius().V1alpha1().Machines())
		ctx.GO(func() {
			factory.Start(ctx.Done())
		})

		// TODO: The number of threads is currently hardcode and needs to be written to the configuration file
		if err = machineController.Run(ctx, opt.ThreadSize); err != nil {
			klog.Fatalln("Error running controller:", err.Error())
			return err
		}

		return nil
	}
}
