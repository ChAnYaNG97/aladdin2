package solvers

import (
	"fmt"
	"k8s.io/kubernetes/schd/aladdin/cores"
	"testing"
)


func TestMaxFlowSolver(t *testing.T) {
	graph := cores.RandomGraphExample()
	source, _ := graph.GetVertex("source")
	sink, _ := graph.GetVertex("sink")
	//graph.PrintGragh()

	solver := NewSMaxFlowSolver(graph, *source, *sink, NewDijkstra())

	flow := solver.MaxFlow2()
	count := 0
	for _, edge := range source.GetOutEdges() {
		if edge.GetCapacity().IsNull() {
			count++
		}
	}
	fmt.Println("布置的任务数：", count)
	fmt.Println("总任务数：", len(source.GetOutEdges()))
	fmt.Println("布置任务的总流量", flow.GetFlow())



	maxCap := cores.NewCpuMemCapacity(0,0)

	for _, ee := range sink.GetInEdges() {
		maxCap.Add(ee.GetMaxCapacity())
	}
	fmt.Println("节点的所有可用流量", maxCap)


	taskCap := cores.NewCpuMemCapacity(0,0)
	for _, ee := range source.GetOutEdges() {
		taskCap.Add(ee.GetMaxCapacity())
	}
	fmt.Println("所有任务的总流量", taskCap)
	//paths := flow.GetPaths()
	//for i := range paths {
	//	fmt.Println(paths[i].GoString())
	//}


	//foe, err := source.GetOutEdge("source-task-0")
	//if err != nil {
	//	fmt.Println("No edge")
	//}
	//
	//ee, err := foe.GetTo().GetOutEdge("task-0-cluster-2")
	//for _, e := range ee.GetTo().GetOutEdges() {
	//	fmt.Println(e.GetCapacity())
	//	fmt.Println(e.GetMaxCapacity())
	//}
}


