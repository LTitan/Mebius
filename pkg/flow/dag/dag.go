package dag

import (
	"fmt"
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
}

type Dag struct {
	head         []*DagNode
	length       int
	raw          map[string]*DagNode
	indegree     map[string]int
	indgreeMutex sync.Mutex
}

// TODO: impl shell and http executor
func (dag *DagNode) Execute() error {
	fmt.Println(dag.Name, "is executing!")
	return nil
}

func (dag *DagNode) ConditionValid() bool {
	return true
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
