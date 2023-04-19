package dag

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/LTitan/Mebius/pkg/utils/function"
	"k8s.io/klog/v2"
)

type DagType string

const (
	shellType   DagType = "shell"
	httpRequest DagType = "http"
)

type DagContext map[string]any

type DagCondition struct {
	Input      string
	Expression string
	Expected   string
}

type DagNode struct {
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Command       string            `json:"command,omitempty"`
	Content       string            `json:"content,omitempty"`
	OutputAlias   map[string]string `json:"output_alias,omitempty"` // alias any=$(echo 'hello')
	Headers       map[string]string `json:"headers,omitempty"`      // for http type
	Timeout       int               `json:"timeout,omitempty"`
	Depends       []string          `json:"depends,omitempty"` // dep and condition
	Conditions    []DagCondition    `json:"conditions,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	Retry         int               `json:"retry,omitempty"`
	RetryWaitTime int               `json:"retry_wait_time,omitempty"`

	// for more context
	PreContext DagContext `json:"-"`
	Next       []*DagNode `json:"-"`
	// utils
	done bool
	mux  sync.Mutex
	fl   function.FunctionLinkInterface
}

type Dag struct {
	head     []*DagNode
	length   int
	finished int
	raw      map[string]*DagNode
	indegree map[string]int
}

func (dag *DagNode) Execute() error {
	var (
		executor Executor
		result   string
	)
	dag.fl = function.NewFunctionLinkErr(
		dag.parseTemplate,
		func() error {
			if dag.Debug {
				klog.Infof("Task [%s] parse dag.Command: %s", dag.Name, dag.Command)
				klog.Infof("Task [%s] parse dag.Content: %s", dag.Name, dag.Content)
			}
			return nil
		},
		func() (err error) {
			// new executor
			switch dag.Type {
			case string(httpRequest):
				executor, err = NewHTTPExecutor(dag.Command, dag.Timeout, dag.Headers)
			case string(shellType):
				executor = &BashExecutor{}
			}
			if err != nil {
				return err
			}
			result, err = executor.Execute(dag.Content)
			return err
		},
	)
	if err := dag.fl.DoErr(); err != nil {
		return err
	}
	if dag.Debug {
		klog.Infof("Task [%s] is executing success, execute result: %s", dag.Name, result)
	}
	if len(dag.OutputAlias) == 0 {
		return nil
	}
	return dag.backwardPropagateContext(result)
}

func (dag *DagNode) parseTemplate() error {
	var err error
	dag.Command, err = warpMapValueToStr(dag.PreContext, dag.Command)
	if err != nil {
		return err
	}
	dag.Content, err = warpMapValueToStr(dag.PreContext, dag.Content)
	if err != nil {
		return err
	}
	return err
}

func (dag *DagNode) backwardPropagateContext(result string) error {
	var (
		err error
	)
	// parse context to the next
	ctx := dag.PreContext
	switch dag.Type {
	case string(shellType):
		for key := range dag.OutputAlias {
			ctx[key] = result
		}
	case string(httpRequest):
		structed := map[string]any{}
		if err = json.Unmarshal([]byte(result), &structed); err != nil {
			return err
		}
		for key, replaceKey := range dag.OutputAlias {
			key = "." + strings.Trim(key, ".")
			value, err := warpMapValueToStrByJQ(structed, key)
			if err != nil {
				return err
			}
			ctx[replaceKey] = value
		}
	}
	// backward propagate
	for _, next := range dag.Next {
		next.mux.Lock()
		if next.PreContext == nil {
			next.PreContext = ctx
		} else {
			for key, value := range ctx {
				next.PreContext[key] = value
			}
		}
		next.mux.Unlock()
	}
	return nil
}

func (dag *DagNode) ConditionValid() bool {
	flag := true
	defer func() {
		if dag.Debug {
			klog.Infof("Task [%s] conditions determine result is: %v", dag.Name, flag)
		}
	}()
	for _, condition := range dag.Conditions {
		if !flag {
			return flag
		}
		key := lint(condition.Input)
		ok, err := warpCondition(dag.PreContext[key], condition.Expression, condition.Expected)
		if err != nil {
			return false
		}
		flag = flag && ok
	}
	return flag
}

func (dag *DagNode) WaitDependDone(stop <-chan struct{}, raw map[string]*DagNode) {
	if len(dag.Depends) <= 0 {
		return
	}
	for {
		select {
		case <-stop:
			return
		default:
			cnt := 0
			for _, dep := range dag.Depends {
				if raw[dep].done {
					cnt++
				}
			}
			if cnt == len(dag.Depends) {
				return
			}
		}
	}
}
