/*
    Reference: github.com/bugra/kmeans
 */

package kmeans

import (
	"math"
	"math/rand"
	"time"
)

type Vector []float64



type ClusteredVector struct {
	ClusterNum int
	Vector
}

func (lhs Vector) Add(rhs Vector) {
	for i, _ := range lhs {
		lhs[i] += rhs[i]
	}
}


func (lhs Vector) Sub(rhs Vector) {
	for i, _ := range lhs {
		lhs[i] -= rhs[i]
	}

}

func (lhs Vector) Sub2(rhs Vector) Vector {
	res := make(Vector, len(lhs))
	for i, v := range lhs {
		res[i] = v - rhs[i]
	}

	return res
}

func (lhs Vector) Mul(scale float64) {
	for i, _ := range lhs {
		lhs[i] *= scale
	}
}

func (lhs Vector) InnerProduct(rhs Vector) Vector {
	var v Vector = make(Vector, len(lhs))
	for i, _ := range lhs {
		v[i] = lhs[i] * rhs[i]
	}
	return v
}

func (v Vector) Norm() float64 {
	var norm float64
	for i, _ := range v {
		norm += v[i] * v[i]
	}

	return math.Sqrt(norm)
}

func (v Vector) Clear() {
	for i, _ := range v {
		v[i] = 0

	}
}

func seed(vecs []Vector, k int) []Vector{
	seeds := make([]Vector, k)
	length := len(vecs)
	// 0 ~ k - 1， for picking the choosen indices
	idxs := make([]int, length)
	for i := 0; i < length; i++ {
		idxs[i] = i
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < k; i++ {
		idx := rand.Intn(length - i)
		seeds[i] = vecs[idxs[idx]]
		idxs[idx], idxs[length - 1 - i] = idxs[length - 1 - i], idxs[idx]
	}
	return seeds
}
//
//func seed(data []ClusteredObservation, k int, distanceFunction DistanceFunction) []Observation {
//	s := make([]Observation, k)
//	s[0] = data[rand.Intn(len(data))].Observation
//	d2 := make([]float64, len(data))
//	for ii := 1; ii < k; ii++ {
//		var sum float64
//		for jj, p := range data {
//			_, dMin := near(p, s[:ii], distanceFunction)
//			d2[jj] = dMin * dMin
//			sum += d2[jj]
//		}
//		target := rand.Float64() * sum
//		jj := 0
//		for sum = d2[0]; sum < target; sum += d2[jj] {
//			jj++
//		}
//		s[ii] = data[jj].Observation
//	}
//	return s
//}

func smartSeed(vecs []Vector, k int, distanceFunction DistanceFunction) []Vector {
	s := make([]Vector, k)
	s[0] = vecs[rand.Intn(len(vecs))]
	d2 := make([]float64, len(vecs))
	for ii := 1; ii < k; ii++ {
		var sum float64
		for jj, p := range vecs {
			_, dMin := closest(ClusteredVector{0, p}, s[:ii], distanceFunction)
			d2[jj] = dMin * dMin
			sum += d2[jj]
		}
		target := rand.Float64() * sum
		jj := 0
		for sum = d2[0]; sum < target; sum += d2[jj] {
			jj++
		}
		s[ii] = vecs[jj]
	}
	return s

}

// Return the index of the nearest cluster and the distance

func closest(p ClusteredVector, mean []Vector, distanceFunction DistanceFunction) (int, float64){
	indexOfCluster := 0
	minDistance := distanceFunction(p.Vector, mean[0])

	for i := 1; i < len(mean); i++ {
		currentDistace := distanceFunction(p.Vector, mean[i])
		if currentDistace < minDistance {
			indexOfCluster = i
			minDistance = currentDistace
		}
	}

	return indexOfCluster, minDistance
}

func kmeans(vecs []Vector, means []Vector, k int, distanceFunction DistanceFunction, threshold float64) ([]ClusteredVector, []Vector){
	clusteredVectors := make([]ClusteredVector, len(vecs))
	for i, _ := range vecs {
		clusteredVectors[i] = ClusteredVector{0,vecs[i]}
	}


	for ; ; {
		meanLength := make([]int, len(means))
		for i, v := range clusteredVectors {
			closestCluster, _ := closest(v, means, distanceFunction)
			clusteredVectors[i].ClusterNum = closestCluster
			meanLength[closestCluster]++
		}


		lastMean := make([]Vector, len(means))
		for i := range lastMean {
			lastMean[i] = make(Vector, len(means[i]))
			copy(lastMean[i], means[i])
		}


		for i := range means {
			means[i] = make(Vector, len(means[i]))
		}

		for _, v := range clusteredVectors {
			means[v.ClusterNum].Add(v.Vector)
		}

		for i ,v := range means {
			v.Mul(1.0 / float64(meanLength[i]))
		}

		end := true

		for i := range means {
			end = end && distanceFunction(means[i], lastMean[i]) < threshold
		}

		if end {
			return clusteredVectors, means
		}

		copy(lastMean, means)


	}

}


func KMeans(raw [][]float64, k int, distanceFunction DistanceFunction, threshold float64) ([]ClusteredVector, []Vector){
	vecs := make([]Vector, len(raw))
	for i, v := range raw {
		vecs[i] = v
	}

	seeds := smartSeed(vecs, k, distanceFunction)
	clustered, means := kmeans(vecs, seeds, k, distanceFunction, threshold)
	return clustered, means
}




