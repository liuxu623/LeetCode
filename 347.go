package main

import (
	"container/heap"
)

type Num struct {
	num  int
	count int
}

type NumHeap []Num

func (h NumHeap) Len() int {
	return len(h)
}

func (h NumHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h NumHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NumHeap) Push(x interface{}) {
	*h = append(*h, x.(Num))
}

func (h *NumHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}
	h1 := make([]Num,len(m))
	for num, count := range m {
		h1 = append(h1, Num{num, count})
	}
	h2 := NumHeap{}
	for i := 0; i < k; i++ {
		h2 = append(h2, h1[i])
	}
	heap.Init(&h2)
	for i := k; i < len(h1); i++ {
		if h2[0].count < h1[i].count {
			heap.Remove(&h2, 0)
			heap.Push(&h2, h1[i])
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-1-i] = h2[0].num
		heap.Remove(&h2, 0)
	}
	return result
}
