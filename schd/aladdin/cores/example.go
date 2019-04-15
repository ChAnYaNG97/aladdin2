package cores

import (
	"fmt"
	"k8s.io/kubernetes/schd/aladdin/data"
	"k8s.io/kubernetes/schd/aladdin/kmeans"
	"math/rand"
	"sort"
	"strconv"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-4
 *
 **************************************************************************************************************/

/*
 * @see <a href=
 *      "https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm">Solver's
 *      Algorithm (Wikipedia)</a> <br>*
 */
 func ShortestPathExample() *Graph {
	 graph := NewGraph()
	 v1    := NewVertex("1")
	 v2    := NewVertex("2")
	 v3    := NewVertex("3")
	 v4    := NewVertex("4")
	 v5    := NewVertex("5")
	 v6    := NewVertex("6")

	 graph.AddVertex(v1)
	 graph.AddVertex(v2)
	 graph.AddVertex(v3)
	 graph.AddVertex(v4)
	 graph.AddVertex(v5)
	 graph.AddVertex(v6)

	 e16   := NewCostEdge(14,NewIntCapacity(1), v1, v6)
	 e12   := NewCostEdge(7, NewIntCapacity(1), v1, v2)
	 e13   := NewCostEdge(9, NewIntCapacity(1), v1, v3)
	 e23   := NewCostEdge(10,NewIntCapacity(1), v2, v3)
	 e24   := NewCostEdge(15,NewIntCapacity(1), v2, v4)
	 e34   := NewCostEdge(11,NewIntCapacity(1), v3, v4)
	 e36   := NewCostEdge(2, NewIntCapacity(1), v3, v6)
	 e45   := NewCostEdge(6, NewIntCapacity(1), v4, v5)
	 e65   := NewCostEdge(9, NewIntCapacity(1), v6, v5)

	 graph.AddEdge(e12)
	 graph.AddEdge(e13)
	 graph.AddEdge(e16)
	 graph.AddEdge(e23)
	 graph.AddEdge(e24)
	 graph.AddEdge(e34)
	 graph.AddEdge(e36)
	 graph.AddEdge(e45)
	 graph.AddEdge(e65)

	 return graph
 }

/*
* @see <a href=
*      "https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm>Solver's
*      Algorithm (Wikipedia)</a> <br>*
*/
func MaxFlowExample() *Graph {
	graph := NewGraph()
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")

	graph.AddVertex(a)
	graph.AddVertex(b)
	graph.AddVertex(c)
	graph.AddVertex(d)
	graph.AddVertex(e)
	graph.AddVertex(f)
	graph.AddVertex(g)


	ab := NewCostEdge(3,NewIntCapacity(1), a, b)
	ad := NewCostEdge(3,NewIntCapacity(1), a, d)
	bc := NewCostEdge(4,NewIntCapacity(1), b, c)
	ca := NewCostEdge(3,NewIntCapacity(1), c, a)
	cd := NewCostEdge(1,NewIntCapacity(1), c, d)
	ce := NewCostEdge(2,NewIntCapacity(1), c, e)
	de := NewCostEdge(2,NewIntCapacity(1), d, e)
	df := NewCostEdge(6,NewIntCapacity(1), d, f)
	eb := NewCostEdge(1,NewIntCapacity(1), e, b)
	eg := NewCostEdge(1,NewIntCapacity(1), e, g)
	fg := NewCostEdge(9,NewIntCapacity(1), f, g)

	graph.AddEdge(ab)
	graph.AddEdge(bc)
	graph.AddEdge(ad)
	graph.AddEdge(ca)
	graph.AddEdge(cd)
	graph.AddEdge(ce)
	graph.AddEdge(de)
	graph.AddEdge(df)
	graph.AddEdge(eb)
	graph.AddEdge(eg)
	graph.AddEdge(fg)

	return graph
}

func CpuMemCapacityExample() *Graph {
	graph := NewGraph()
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")
	h := NewVertex("H")
	i := NewVertex("I")

	graph.AddVertex(a)
	graph.AddVertex(b)
	graph.AddVertex(c)
	graph.AddVertex(d)
	graph.AddVertex(e)
	graph.AddVertex(f)
	graph.AddVertex(g)
	graph.AddVertex(h)
	graph.AddVertex(i)


	ab := NewCostEdge(1, NewCpuMemCapacity(8,8), a, b)
	ac := NewCostEdge(1, NewCpuMemCapacity(4,8), a, c)
	ad := NewCostEdge(1, NewCpuMemCapacity(4,8), a, d)

	be := NewCostEdge(1, NewCpuMemCapacity(8,8), b, e)
	bf := NewCostEdge(10, NewCpuMemCapacity(8,8), b, f)

	ce := NewCostEdge(10, NewCpuMemCapacity(4,8), c, e)
	cf := NewCostEdge(1, NewCpuMemCapacity(4,8), c, f)


	de := NewCostEdge(10, NewCpuMemCapacity(4,8), d, e)
	df := NewCostEdge(1, NewCpuMemCapacity(4,8), d, f)

	eg := NewCostEdge(1, NewCpuMemCapacity(8,8), e, g)
	fh := NewCostEdge(1, NewCpuMemCapacity(8, 16), f, h)

	gi := NewCostEdge(1, NewCpuMemCapacity(8,8), g, i)
	hi := NewCostEdge(1, NewCpuMemCapacity(8, 16), h, i)

	graph.AddEdge(ab)
	graph.AddEdge(ac)
	graph.AddEdge(ad)

	graph.AddEdge(be)
	graph.AddEdge(bf)

	graph.AddEdge(ce)
	graph.AddEdge(cf)

	graph.AddEdge(de)
	graph.AddEdge(df)

	graph.AddEdge(eg)
	graph.AddEdge(fh)

	graph.AddEdge(gi)
	graph.AddEdge(hi)

	return graph

}

const (
	TaskNum = 34594
	NodeNum = 4338

)

func RandomGraphExample() *Graph {
	graph := NewGraph()
	source := NewVertex("source")
	sink := NewVertex("sink")
	graph.AddVertex(source)
	graph.AddVertex(sink)

	rand.Seed(42)



	tasks, nodes := data.ReadDataFromJsonFile()

	// 随机生成任务节点，和从source到他们的边
	for i := 0; i < TaskNum; i++ {
		v := NewVertex("task-" + strconv.Itoa(i))
		e := NewCostEdge(1, NewCpuMemCapacity(tasks[i].Cpu, tasks[i].Mem), source, v)
		graph.AddVertex(v)
		graph.AddEdge(e)
	}
	// 生成机器节点，和他们到sink的边
	for i := 0; i < NodeNum; i++ {
		v := NewVertex("node-" + strconv.Itoa(i))
		e := NewCostEdge(1, NewCpuMemCapacity(nodes[i].Cpu, nodes[i].Mem), v, sink)
		graph.AddVertex(v)
		graph.AddEdge(e)
		}


	//生成聚合节点
	var sortedNames []string
	for toSinkEdgeName, _ := range sink.GetInEdges() {
		sortedNames = append(sortedNames, toSinkEdgeName)
	}

	sort.Strings(sortedNames)
	rawData := make([][]float64, len(sortedNames))
	for i, name := range sortedNames {
		ee, _:= sink.GetInEdge(name)

		cap := ee.capacity.(*CpuMemCapacity)
		rawData[i] = []float64{float64(cap.mem) / float64(cap.cpu) / 1000}
	}

	fmt.Println("End Raw Data")
	labels, means := kmeans.KMeans(rawData, 9, kmeans.EuclideanDistanceFunction, 0.001)
	fmt.Println(labels)
	fmt.Println(means)

	for i := 0; i < len(means); i++ {
		v := NewVertex("cluster-" + strconv.Itoa(i))
		graph.AddVertex(v)
	}

	for i, name := range sortedNames {
		label := labels[i].ClusterNum
		ee, _ := sink.GetInEdge(name)
		vv := ee.GetFrom()
		cv, _ := graph.GetVertex("cluster-" + strconv.Itoa(label))
		ce := NewCostEdge(1, NewCpuMemCapacityWithCapacity(ee.GetCapacity()), cv, vv)
		graph.AddEdge(ce)
	}


	for _, ee := range source.GetOutEdges() {
		var ids kmeans.IndexDistanceSlice
		v1 := ee.GetTo()
		v1Cap := ee.capacity.(*CpuMemCapacity)
		v1Ratio := kmeans.Vector{float64(v1Cap.mem) / float64(v1Cap.cpu) / 1000}
		for i := 0; i < len(means); i++ {
			ids = append(ids, kmeans.IndexDistance{i, kmeans.EuclideanDistanceFunction(v1Ratio, means[i])})
		}
		sort.Sort(ids)
		coff := 1
		for i := 0; i < len(ids); i++ {
			cv, _ := graph.GetVertex("cluster-" + strconv.Itoa(ids[i].Index))
			graph.AddEdge(NewCostEdge(coff, NewCpuMemCapacityWithCapacity(v1Cap), v1, cv))
			coff *= 2
		}
	}





	return graph


	// 任务到聚合的边

	// 聚合到机器的边
}