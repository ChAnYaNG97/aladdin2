package cores

import (
	"fmt"
	"sort"
	"testing"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 *
 **************************************************************************************************************/

 func TestNewGraph(t *testing.T) {
	 graph := ShortestPathExample()

	 _, vv := graph.GetVertex("1")

	 if vv != nil {
	 	t.Error();
	 }

	 _, nvv := graph.GetVertex("7")

	 if nvv == nil {
	 	t.Error()
	 }

	 _, ve := graph.GetEdge("1-2")

	 if ve != nil {
		 t.Error();
	 }

	 _, nve := graph.GetEdge("1-7")

	 if nve == nil {
		 t.Error()
	 }
 }


 func TestCpuMemGraph(t *testing.T) {
 	graph := CpuMemCapacityExample()

 	graph.PrintGragh()
 	source, _ := graph.GetVertex("A")
 	sink, _ := graph.GetVertex("I")
 	//solver := solvers.NewSMaxFlowSolver(graph, *source, *sink, solvers.NewDijkstra())

 	for _, edge := range source.GetOutEdges() {
 		fmt.Println(edge)

	}
 	for _, edge := range sink.GetInEdges() {
 		fmt.Println(edge)
	}

 	b, _ := graph.GetVertex("B")

 	fmt.Println()
 	for _, edge := range b.GetOutEdges() {
 		fmt.Println(edge)
	}

 }


 func TestSortEdges(t *testing.T) {
 	es := make(EdgeSlice, 0)
 	for i := 0; i < 4; i++ {
 		ee := NewCostEdge((4 - i), NewCpuMemCapacity(1,1), nil, nil)

 		es = append(es,*ee)
 		}
 	sort.Sort(es)
 	for _, ee := range es {
 		fmt.Println(ee.GetCost())
	}

 }