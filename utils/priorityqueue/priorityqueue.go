package priorityqueue

import (
	"container/heap"
)

// Item is something that can be put into priority queue
type Item struct {
	Value   string
	Priority int
	index   int
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

// Len returns the number of items in the queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less returns true if item at index j is less than item at index j
func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swaps items at indexes i and j
func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push inserts an item into the queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes an item form the queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Peek returns the item, without removing it from the queue, if present, otherwise nil
func (pq PriorityQueue) Peek(value string) interface{} {
	for i := range pq {
		if pq[i].Value == value {
			return pq[i]
		}
	}
	return nil
}

// Update changes the priority of an item
func (pq *PriorityQueue) Update(value string, priority int) {
	for i := range *pq {
		if (*pq)[i].Value == value {
			(*pq)[i].Priority = priority
			heap.Fix(pq, (*pq)[i].index)
		}
	}
}

// Has returns true if the queue contains the value
func (pq PriorityQueue) Has(value string) bool {
	for i := range pq {
		if pq[i].Value == value {
			return true
		}
	}
	return false
}
