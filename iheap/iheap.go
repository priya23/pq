package iheap

import (
	"container/heap"
)

type Item struct {
	Val      interface{}
	Priority int
	//	Index    int
}

type PriorityQueue []*Item

//for sort implementation
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//implementing default interface
func (pq *PriorityQueue) Push(x interface{}) {
	//	n := len(*pq)
	item := x.(*Item)
	//	item.Index = n
	*pq = append(*pq, item)

}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func CreateNew(val string, priority int) *Item {
	itt := Item{Val: val, Priority: priority}
	return &itt
}

func CreateHeap() *PriorityQueue {
	p := make(PriorityQueue, 0)
	heap.Init(&p)
	return &p
}
