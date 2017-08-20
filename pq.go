package pq

import (
	//"container/heap"
	"github.com/priya23/pq/iheap"
	//	"fmt"
)

type PQ interface {
	Push(k interface{})
	Pop() rItem
}

func NewQueue() PQ {
	return iheap.CreateHeap()
}

func CreateNode(val string, priority int) interface{} {
	return iheap.CreateNew(val, priority)
}

type Value interface {
	Get() interface{}
}
