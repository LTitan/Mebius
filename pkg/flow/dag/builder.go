package dag

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func BuildDag(nodes []*DagNode) *Dag {
	dag := &Dag{
		indegree: make(map[string]int, len(nodes)),
		length:   len(nodes),
		raw:      make(map[string]*DagNode, len(nodes)),
	}
	for i, node := range nodes {
		dag.raw[node.Name] = nodes[i]
		dag.indegree[node.Name] = len(node.Depends)
	}
	for name, indegree := range dag.indegree {
		if indegree == 0 {
			dag.raw[name].PreContext = make(DagContext)
			dag.head = append(dag.head, dag.raw[name])
		}
		for _, dependName := range dag.raw[name].Depends {
			dag.raw[dependName].Next = append(dag.raw[dependName].Next,
				dag.raw[name],
			)
		}
	}
	return dag
}

func (dag *Dag) Run(ctx context.Context) error {
	return dag.executeIndgreeZeroNodes(ctx, dag.head)
}

func (dag *Dag) executeIndgreeZeroNodes(ctx context.Context, nodes []*DagNode) error {
	if len(nodes) == 0 {
		return nil
	}
	if err := dag.executeNodes(ctx, nodes); err != nil {
		return err
	}
	tmpNodes := []*DagNode{}
	for _, head := range nodes {
		for _, next := range head.Next {
			dag.indegree[next.Name]--
			if dag.indegree[next.Name] == 0 {
				tmpNodes = append(tmpNodes, next)
			}
		}
	}
	return dag.executeIndgreeZeroNodes(ctx, tmpNodes)
}

func (dag *Dag) executeNodes(ctx context.Context, nodes []*DagNode) error {
	wg := errgroup.Group{}
	exec := func(nd *DagNode) func() error {
		return func() error {
			defer func() {
				nd.done = true
			}()
			// wait depends
			nd.WaitDependDone(ctx.Done(), dag.raw)
			// execute
			if !nd.ConditionValid() {
				// not execute
				return nil
			}
			return nd.Execute()
		}
	}
	for _, node := range nodes {
		wg.Go(exec(node))
	}
	return wg.Wait()
}
