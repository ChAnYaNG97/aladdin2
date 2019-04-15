package solvers

import (
	"k8s.io/kubernetes/schd/aladdin/cores"
	"sort"
)

func (s *Solver) MaxFlow2() cores.Flow {
	flow := cores.NewFlow()


	for {
		path := s.AvailablePath3()



		if path.GetCost() == -1 {
			break
		}

		minFlow := s.graph.UpdateGraghForMaxFlow2(path)

		flow.AddFlow(minFlow)
		flow.AddPath(path)
		//s.graph.PrintGragh()



	}
	return flow
}




func (s *Solver) AvailablePath2() cores.Path {
	for _, firstOutEdge := range s.source.GetOutEdges() {
		if firstOutEdge.GetCapacity().IsNull() {
			continue
		}

		// 未调度的任务节点
		firstLevelVertex := firstOutEdge.GetTo()

		minCostEdgeName := ""
		var minCostEdge cores.Edge
		//secondOutEdges := make(cores.EdgeSlice,0)
		//
		//for _, secondOutEdge := range firstLevelVertex.GetOutEdges() {
		//	secondOutEdges = append(secondOutEdges, secondOutEdge)
		//}
		//
		//sort.Sort(secondOutEdges)





		for secondEdgeName, secondOutEdge := range firstLevelVertex.GetOutEdges() {

			// 找到 cost 最小的分类节点对应的边
			if minCostEdgeName == "" {
				minCostEdgeName = secondEdgeName
				minCostEdge = secondOutEdge
			}


			if secondOutEdge.GetCost() < minCostEdge.GetCost() {
				minCostEdgeName = secondEdgeName
				minCostEdge = secondOutEdge
			}

		}

		aggregatorNode := minCostEdge.GetTo()
		for _, thirdEdge := range aggregatorNode.GetOutEdges() {
			// 找到所有的最小cost的聚合节点的对应的机器
			thirdLevelVertex := thirdEdge.GetTo()
			for _, toSinkEdge := range thirdLevelVertex.GetOutEdges() {
				cmCap := toSinkEdge.GetCapacity().(*cores.CpuMemCapacity)
				if cmCap.GreatEqual(firstOutEdge.GetCapacity()) {
					currentPath := cores.NewPath()
					currentPath.AddEdge(firstOutEdge)
					currentPath.AddEdge(minCostEdge)
					currentPath.AddEdge(thirdEdge)
					currentPath.AddEdge(toSinkEdge)
					return *currentPath
				}
			}
		}
	}












	//for _, firstOutEdge := range s.source.GetOutEdges(){
	//	// 找到所有未调度的任务节点
	//	if firstOutEdge.GetCapacity().IsNull() {
	//		continue
	//	}
	//	firstLevelVertex := firstOutEdge.GetTo()
	//	for secondEdgeName, secondOutEdge := range firstLevelVertex.GetOutEdges(){
	//		// 找到所有的机器节点
	//		secondLevelVertex := secondOutEdge.GetTo()
	//		// 是否可以直接部署到该机器
	//		capacityNeed := cores.NewIntCapacity(0)  // 如果不能部署，需要反悔的最少容量；最大容量是它本身
	//		for _, toSinkEdge := range secondLevelVertex.GetOutEdges(){
	//			// 如果可以直接部署
	//			if !toSinkEdge.GetCapacity().Less(firstOutEdge.GetCapacity()) {
	//				currentPath := cores.NewPath()
	//				currentPath.AddEdge(firstOutEdge)
	//				currentPath.AddEdge(secondOutEdge)
	//				currentPath.AddEdge(toSinkEdge)
	//				return *currentPath
	//			}else {
	//
	//				tmpCapacity, err := firstOutEdge.GetCapacity().Sub2(toSinkEdge.GetCapacity())
	//				if err != nil {
	//					panic("this step shouldn't excute")
	//				}
	//				capacityNeed.Add(tmpCapacity)
	//				fmt.Println("schedule pod "+firstLevelVertex.GetName()+" to node "+secondLevelVertex.GetName()+" need other "+capacityNeed.GoString())
	//			}
	//		}
	//		// 不能直接部署，就需要考虑反悔哪一个任务
	//		var selectedEdge cores.Edge
	//		var selectedEdgeName = ""
	//
	//		for thirdEdgeName, thirdInEdge := range secondLevelVertex.GetInEdges(){
	//			// 反悔任务时本任务，跳过
	//			if thirdEdgeName == secondEdgeName{
	//				continue
	//			}
	//			// 反悔任务流量小于最小容量需求 或者 不小于任务本身的容量需求，跳过
	//
	//			fmt.Println("---开始获得反向边容量---")
	//			fmt.Println("Edge's maxCapacity : "+thirdInEdge.GetMaxCapacity().GoString())
	//			fmt.Println("Edge's capacity : "+thirdInEdge.GetCapacity().GoString())
	//			fmt.Println("---结束获得反向边容量---")
	//
	//			fmt.Println("Selecting a task to prrmpt: this edge "+thirdInEdge.GetName()+"'s reverseEdge has "+thirdInEdge.GetReverseCapacity().GoString())
	//			fmt.Println("Selecting a task to prrmpt: this pod needs "+firstOutEdge.GetCapacity().GoString())
	//
	//			if thirdInEdge.GetReverseCapacity().Less(capacityNeed) || !thirdInEdge.GetReverseCapacity().Less(firstOutEdge.GetCapacity()){
	//				continue
	//			}
	//			// 找到大于最小容量需求的最小任务进行反悔
	//			if selectedEdgeName == "" {
	//				selectedEdge = thirdInEdge
	//				selectedEdgeName = thirdEdgeName
	//				continue
	//			}
	//			if thirdInEdge.GetReverseCapacity().Less(selectedEdge.GetReverseCapacity()) {
	//				selectedEdge = thirdInEdge
	//				selectedEdgeName = thirdEdgeName
	//			}
	//
	//		}
	//		// 如果找不到可反悔的任务，则直接返回无可行路径
	//		if selectedEdgeName == ""{
	//			fmt.Println("Can't find a task to preempt")
	//			break
	//		}
	//
	//		thirdLevelVertex := selectedEdge.GetFrom()
	//		// 反悔的任务是否可以放在其他的机器上
	//		for fourthEdgeName, fourthOutEdge := range thirdLevelVertex.GetOutEdges(){
	//			if fourthEdgeName == selectedEdgeName {
	//				continue
	//			}
	//			fourthLevelVertex := fourthOutEdge.GetTo() // 获得待反悔的机器节点
	//			// 判断该待反悔的机器节点是否有足够的容量
	//			for _, toSinkEdge := range fourthLevelVertex.GetOutEdges(){
	//				// 如果可以反悔到这个机器上
	//				if !toSinkEdge.GetCapacity().Less(fourthOutEdge.GetCapacity()) {
	//					currentPath := cores.NewPath()
	//					currentPath.AddEdge(firstOutEdge)
	//					currentPath.AddEdge(secondOutEdge)
	//					currentPath.AddEdge(selectedEdge)
	//					currentPath.AddEdge(fourthOutEdge)
	//					currentPath.AddEdge(toSinkEdge)
	//					return *currentPath
	//				}
	//			}
	//
	//		}
	//		// 如果不能反悔到别的机器上，就直接反悔整个任务，返回的任务不部署
	//		for _, fromSourceEdge := range thirdLevelVertex.GetInEdges(){
	//			currentPath := cores.NewPath()
	//			currentPath.AddEdge(firstOutEdge)
	//			currentPath.AddEdge(secondOutEdge)
	//			currentPath.AddEdge(selectedEdge)
	//			currentPath.AddEdge(fromSourceEdge)
	//			return *currentPath
	//		}
	//
	//
	//	}
	//}

	return *cores.NewInvalidPath()
}



func (s *Solver) AvailablePath3() cores.Path {
	for _, firstOutEdge := range s.source.GetOutEdges() {
		if firstOutEdge.GetCapacity().IsNull() {
			continue
		}

		// 未调度的任务节点
		firstLevelVertex := firstOutEdge.GetTo()

		secondOutEdges := make(cores.EdgeSlice,len(firstLevelVertex.GetOutEdges()))
		// secondOutEdges := make(cores.PriorityQueue,len(firstLevelVertex.GetOutEdges()))
		i := 0
		for _, secondOutEdge := range firstLevelVertex.GetOutEdges() {
			secondOutEdges[i] = secondOutEdge
			i++
		}


		// 构建一个最小堆
		//heap.Init(&secondOutEdges)
		sort.Sort(secondOutEdges)



		for _, secondOutEdge := range secondOutEdges {
			//secondOutEdge := secondOutEdges.Pop().(cores.Edge)
			aggregatorNode := secondOutEdge.GetTo()

			for _, thirdEdge := range aggregatorNode.GetOutEdges() {
				// 找到所有的最小cost的聚合节点的对应的机器
				thirdLevelVertex := thirdEdge.GetTo()
				for _, toSinkEdge := range thirdLevelVertex.GetOutEdges() {
					cmCap := toSinkEdge.GetCapacity().(*cores.CpuMemCapacity)
					if cmCap.GreatEqual(firstOutEdge.GetCapacity()) {
						//fmt.Println("Cost is ", secondOutEdge.GetCost())
						//fmt.Println("--------------------")
						currentPath := cores.NewPath()
						currentPath.AddEdge(firstOutEdge)
						currentPath.AddEdge(secondOutEdge)
						currentPath.AddEdge(thirdEdge)
						currentPath.AddEdge(toSinkEdge)
						return *currentPath
					}
				}
			}

			// 找不到一台机器去放，需要反悔一些（？）任务


		}


	}

	return *cores.NewInvalidPath()
}