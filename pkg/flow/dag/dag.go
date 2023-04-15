package dag

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
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
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Command string `json:"command,omitempty"`
	Content string `json:"content,omitempty"`
	// alias any=$(echo 'hello')
	OutputAlias map[string]string `json:"output_alias,omitempty"`
	// for http type
	Headers map[string]string `json:"headers,omitempty"`
	Timeout int               `json:"timeout,omitempty"`
	// dep and condition
	Depends    []string       `json:"depends,omitempty"`
	Conditions []DagCondition `json:"conditions,omitempty"`

	// for more context
	PreContext DagContext `json:"-"`
	Next       []*DagNode `json:"-"`

	done bool
	mux  sync.Mutex
}

type Dag struct {
	head     []*DagNode
	length   int
	raw      map[string]*DagNode
	indegree map[string]int
}

// TODO: impl shell and http executor
func (dag *DagNode) Execute() error {
	var (
		executor Executor
		err      error
	)
	dag.Command, err = warpMapValueToStr(dag.PreContext, dag.Command)
	if err != nil {
		return err
	}
	dag.Content, err = warpMapValueToStr(dag.PreContext, dag.Content)
	if err != nil {
		return err
	}
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
	result, err := executor.Execute(dag.Content)
	if err != nil {
		return err
	}
	fmt.Println(dag.Name, "is executing success!")
	if len(dag.OutputAlias) == 0 {
		return nil
	}

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
