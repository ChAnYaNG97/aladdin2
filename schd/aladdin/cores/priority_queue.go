package cores




type PriorityQueue []Edge




func (pq PriorityQueue) Len() int{
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].GetCost() > pq[j].GetCost()
	// TODO: !Less() 和 GreatEqual()
}


func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	edge := x.(Edge)
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}