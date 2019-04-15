package cores

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	edges := PriorityQueue{
		Edge{cost:1},
		Edge{cost:34},
		Edge{cost:2},
		Edge{cost:8},
	}


	heap.Init(&edges)

	for edges.Len() > 0 {
		fmt.Println(edges.Pop().(Edge))
	}
}