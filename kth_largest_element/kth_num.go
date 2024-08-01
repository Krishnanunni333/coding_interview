package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n - 1]
	return x
}

func main(){
	arr := &MaxHeap{4,2,9,7,5,6,7,1,3}
	target := 4
	heap.Init(arr)

	// heap.Push(arr, 93)
	for i := 1; i < target; i++{
		heap.Pop(arr)
	}
	min := heap.Pop(arr).(int)
	fmt.Println("MaxHeap after heap.Pop():", *arr)
	fmt.Println("Popped element:", min)
}