package main

import (
	"container/heap"
	"fmt"
)

//at first,we should define the type of our heap
//type IntHeap []int
//带权值的形式
type IntHeap [][2]int

//then implement the following 5 methods in order to use heap
func (h IntHeap) Len() int { return len(h) }

//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
//带权值的less
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//func (h *IntHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//带权值的push
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//347
func topKFrequent(nums []int, k int) []int {
	dic := map[int]int{}
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		dic[v]++
	}
	for _, v := range nums {
		//带权值的push
		heap.Push(h, [2]int{dic[v], v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := []int{}
	for i := 0; i < k; i++ {
		//.(int)方法将interface转int
		//ans = append(ans, heap.Pop(h).(int))

		//带权值的pop
		ans = append(ans, heap.Pop(h).([2]int)[1])
	}
	return ans
}


