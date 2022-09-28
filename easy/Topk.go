package easy

import (
	"container/heap"
)

// TopK 最大K值
func Topk(arr []int, k int) (res []int) {
	h := new(myHeap)
	for i := 0; i < k; i++ {
		h.Push(arr[i])
	}
	heap.Init(h)
	for i := k + 1; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	return
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
