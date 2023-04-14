package factory

import (
	"sync"
	"time"

	mebiusclientset "github.com/LTitan/Mebius/pkg/clients/clientset/mebius"
	"github.com/LTitan/Mebius/pkg/clients/informer/externalversions"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/utils/function"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	once = sync.Once{}
)

type Application interface {
	RegisterCommand()
}

type FrameworkInterface interface {
	Init() error
	GetKubeConfig() *rest.Config
	GetMebiusSharedInformerFactory() externalversions.SharedInformerFactory
	GetMebiusClientSet() mebiusclientset.Interface
	GetKubeClientSet() kubernetes.Interface
}

// some base components
type BaseFramework struct {
	opts                        *options.GlobalOption
	kubeConfig                  *rest.Config
	mebiusSharedInformerFactory externalversions.SharedInformerFactory
	mebiusClientSet             mebiusclientset.Interface
	kubeClientSet               kubernetes.Interface
}

func NewBaseFramework(opts *options.GlobalOption) FrameworkInterface {
	return &BaseFramework{
		opts: opts,
	}
}

func (bf *BaseFramework) Init() error {
	// only init once
	var err error
	once.Do(
		func() {
			err = function.NewFunctionLinkErr(
				bf.buildKubeConfig,
				bf.buildKubeClientSet,
				bf.buildClientSet,
				bf.buildInformer).DoErr()
		},
	)
	return err
}

func (bf *BaseFramework) GetKubeConfig() *rest.Config {
	return bf.kubeConfig
}

func (bf *BaseFramework) GetMebiusSharedInformerFactory() externalversions.SharedInformerFactory {
	return bf.mebiusSharedInformerFactory
}

func (bf *BaseFramework) GetMebiusClientSet() mebiusclientset.Interface {
	return bf.mebiusClientSet
}

func (bf *BaseFramework) GetKubeClientSet() kubernetes.Interface {
	return bf.kubeClientSet
}

func (bf *BaseFramework) buildKubeConfig() (err error) {
	bf.kubeConfig, err = clientcmd.BuildConfigFromFlags("", bf.opts.KubeConfig)
	return
}

func (bf *BaseFramework) buildClientSet() (err error) {
	bf.mebiusClientSet, err = mebiusclientset.NewForConfig(bf.kubeConfig)
	return
}

func (bf *BaseFramework) buildKubeClientSet() (err error) {
	bf.kubeClientSet, err = kubernetes.NewForConfig(bf.kubeConfig)
	return
}

func (bf *BaseFramework) buildInformer() (err error) {
	bf.mebiusSharedInformerFactory = externalversions.NewSharedInformerFactory(bf.mebiusClientSet,
		time.Hour*time.Duration(bf.opts.Controller().ResyncPeriod))
	return
}
