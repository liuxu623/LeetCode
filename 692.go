package main

import (
	"container/heap"
)

type Word struct {
	word  string
	count int
}

type WordHeap []Word

func (h WordHeap) Len() int {
	return len(h)
}

func (h WordHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].word > h[j].word
	}
	return h[i].count < h[j].count
}

func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(Word))
}

func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	h1 := make([]Word,len(m))
	for word, count := range m {
		h1 = append(h1, Word{word, count})
	}
	h2 := WordHeap{}
	for i := 0; i < k; i++ {
		h2 = append(h2, h1[i])
	}
	heap.Init(&h2)
	for i := k; i < len(h1); i++ {
		if h2[0].count < h1[i].count || (h2[0].count == h1[i].count && h2[0].word > h1[i].word) {
			heap.Remove(&h2, 0)
			heap.Push(&h2, h1[i])
		}
	}
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[k-1-i] = h2[0].word
		heap.Remove(&h2, 0)
	}
	return result
}
