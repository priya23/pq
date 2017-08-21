package pq

import (
	//"container/heap"
	"github.com/priya23/pq/iheap"
	//	"fmt"
)

type PQ interface {
	Push(k interface{})
	Pop() interface{}
	Update(interface{}, interface{}, int)
}

func NewQueue() PQ {
	return iheap.CreateHeap()
}

func CreateNode(val interface{}, priority int) interface{} {
	return iheap.CreateNew(val, priority)
}
