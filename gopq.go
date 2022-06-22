package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{make(map[P][]V), priorities}
}

func (pq *PriorityQueue[P, V]) add(priority P, val V) {
	pq.items[priority] = append(pq.items[priority], val)
}

func (pq *PriorityQueue[P, V]) next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		p := pq.priorities[i]
		items := pq.items[p]
		if len(items) > 0 {
			next := items[0]
			pq.items[p] = items[1:]
			return next, true
		}
	}
	var empty V
	return empty, false
}

func main() {
	pq := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	pq.add(Low, "L-1")
	pq.add(High, "H-1")

	fmt.Println(pq.next())

	pq.add(Medium, "M-1")
	pq.add(High, "H-2")
	pq.add(High, "H-3")

	fmt.Println(pq.next())
	fmt.Println(pq.next())
	fmt.Println(pq.next())
	fmt.Println(pq.next())
	fmt.Println(pq.next())
}
