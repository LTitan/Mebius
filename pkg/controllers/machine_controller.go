package controllers

import (
	"fmt"
	"time"

	"github.com/LTitan/Mebius/pkg/apis/v1alpha1"
	informers "github.com/LTitan/Mebius/pkg/clients/informer/externalversions/apis/v1alpha1"
	mcontext "github.com/LTitan/Mebius/pkg/context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const (
	machineControllerName = "machine-controller"
)

type MachineController struct {
	informer  informers.MachineInformer
	workqueue workqueue.RateLimitingInterface
}

func NewMachineController(informer informers.MachineInformer) Interface {
	machineController := &MachineController{
		informer:  informer,
		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machine"),
	}

	klog.Infoln("Setting up Machine event handlers")

	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    machineController.addMachine,
		DeleteFunc: machineController.deleteMachine,
		UpdateFunc: machineController.updateMachine,
	})

	return machineController
}

func (c *MachineController) Run(ctx mcontext.MContext, threadSize int) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShuttingDown()

	klog.Infoln("Starting Machine control loop")

	klog.Infoln("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(ctx.Done(), c.informer.Informer().HasSynced); !ok {
		return fmt.Errorf("failed to wati for caches to sync")
	}

	klog.Infoln("Starting workers")
	for i := 0; i < threadSize; i++ {
		go wait.Until(c.runWorker, time.Second, ctx.Done())
	}

	klog.Infoln("Started workers")
	ctx.WaitGoroutine()
	klog.Infoln("Shutting down workers")
	return nil
}

func (c *MachineController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *MachineController) processNextWorkItem() bool {
	item, shutdown := c.workqueue.Get()
	if shutdown {
		return false
	}

	if err := func(item interface{}) error {
		defer c.workqueue.Done(item)
		var (
			key string
			ok  bool
		)
		if key, ok = item.(string); !ok {
			c.workqueue.Forget(item)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", item))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s':%s", item, err.Error())
		}
		c.workqueue.Forget(item)
		return nil
	}(item); err != nil {
		runtime.HandleError(err)
		return false
	}
	return true
}

func (c *MachineController) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid respirce key:%s", key))
	}

	astro, err := c.informer.Lister().Machines(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		runtime.HandleError(fmt.Errorf("failed to get machine by: %s/%s", namespace, name))
		return err
	}
	fmt.Printf("[MachineCRD] try to process machine:%#v ...\n", astro)
	// TODO: do something
	return nil
}

func (c *MachineController) addMachine(item interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(item); err != nil {
		runtime.HandleError(err)
		return
	}

	klog.Infoln("add machine crd")

	c.workqueue.AddRateLimited(key)
}
func (c *MachineController) deleteMachine(item interface{}) {
	var key string
	var err error
	if key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(item); err != nil {
		runtime.HandleError(err)
		return
	}

	klog.Infoln("delete machine crd")

	c.workqueue.AddRateLimited(key)
}
func (c *MachineController) updateMachine(old, new interface{}) {
	oldItem := old.(*v1alpha1.Machine)
	newItem := new.(*v1alpha1.Machine)
	if oldItem.ResourceVersion == newItem.ResourceVersion {
		return
	}
	klog.Infoln("update machine crd")
	c.workqueue.AddRateLimited(new)
}
