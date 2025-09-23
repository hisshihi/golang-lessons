package main

import "fmt"

const (
	High = iota
	Medium
	Low
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{items: map[P][]V{}, priorities: priorities}
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		item := pq.items[priority]
		if len(item) > 0 {
			next := item[0]
			pq.items[priority] = item[1:]
			return next, true
		}
	}
	return *new(V), false
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})

	queue.Add(Low, "L-1")
	queue.Add(Medium, "M-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())
}
