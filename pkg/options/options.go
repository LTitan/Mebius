package options

import (
	"time"

	mebiusclientset "github.com/LTitan/Mebius/pkg/clients/clientset/mebius"
	"github.com/LTitan/Mebius/pkg/clients/informer/externalversions"
	"github.com/LTitan/Mebius/pkg/controllers"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

type GlobalOption struct {
	EnableKubeConfig bool
	KubeConfig       string

	Cmd *cobra.Command
}

func (g *GlobalOption) Parse() {
	g.Cmd.PersistentFlags().BoolVar(&g.EnableKubeConfig, "enable-kubeconfig", true, "is enable kube-config")
	g.Cmd.PersistentFlags().StringVar(&g.KubeConfig, "kubeconfig", clientcmd.RecommendedHomeFile, "kube-config file")
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
			// TODO: We may need a subcommand to execute it
			RunE: runController,
		},
	}
	klog.InitFlags(nil)
	g.Parse()
	return g
}

func runController(cmd *cobra.Command, args []string) error {
	// TODO: Deal with enable-kubeconfig
	kubeconfig := cmd.PersistentFlags().Lookup("kubeconfig").Value.String()
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := setupSignalHandler()

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

	factory := externalversions.NewSharedInformerFactory(clientset, time.Second*30)

	machineController := controllers.NewMachineController(factory.Mebius().V1alpha1().Machines())

	go factory.Start(stopCh)

	// TODO: The number of threads is currently hardcode and needs to be written to the configuration file
	if err = machineController.Run(2, stopCh); err != nil {
		klog.Fatalln("Error running controller:", err.Error())
		return err
	}

	return nil
}
