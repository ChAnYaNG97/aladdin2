package kmeans



type IndexDistance struct {
	Index int
	Distance float64
}

type IndexDistanceSlice []IndexDistance

func (ids IndexDistanceSlice) Len() int{
	return len(ids)
}

func (ids IndexDistanceSlice) Less(i, j int) bool {
	return ids[i].Distance < ids[j].Distance
}

func (ids IndexDistanceSlice) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

type DistanceFunction func(first, second Vector) float64


func EuclideanDistanceFunction(first, second Vector) float64 {
	res := first.Sub2(second)
	return res.Norm()
}




