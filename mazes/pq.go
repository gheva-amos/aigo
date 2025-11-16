package maze

import (
	"container/heap"
)

type PriorityQueue struct {
	heap      point_heap
	direction bool
}

func NewPriorityQueue(direction bool) *PriorityQueue {
	pq := &PriorityQueue{}
	pq.direction = direction
	heap.Init(&pq.heap)
	return pq
}

func (pq *PriorityQueue) Push(value Point, priority int) {
	heap.Push(&pq.heap, &point_item{
		ThePoint:  value,
		Priority:  priority,
		HighFirst: pq.direction,
	})
}

func (pq *PriorityQueue) Pop() Point {
	item := heap.Pop(&pq.heap).(*point_item)
	return item.ThePoint
}

func (pq *PriorityQueue) Len() int {
	return pq.heap.Len()
}

func (pq *PriorityQueue) Find(point Point) bool {
	for _, n := range pq.heap {
		if point.Equals(n.ThePoint) {
			return true
		}
	}
	return false
}

type point_item struct {
	ThePoint  Point
	Priority  int
	index     int
	HighFirst bool
}

type point_heap []*point_item

func (h point_heap) Len() int {
	return len(h)
}

func (h point_heap) Less(i, j int) bool {
	if h[i].HighFirst {
		return h[i].Priority < h[j].Priority
	}
	return h[i].Priority > h[j].Priority
}

func (h point_heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *point_heap) Push(x interface{}) {
	item := x.(*point_item)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *point_heap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[:n-1]
	return item
}
