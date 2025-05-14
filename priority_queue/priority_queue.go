package priorityqueue

import (
	"sort"

	"github.com/DatNguyenPT/go-ds/pair"
)

// PriorityQueue represents a priority queue implemented with Pair elements.
type PriorityQueue struct {
	Data []pair.Pair
}

// byA is a custom type to sort Pair items by field A in ascending order.
// A lower value in field A has a higher priority.
type byA []pair.Pair

func (p byA) Len() int      { return len(p) }
func (p byA) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byA) Less(i, j int) bool {
	// Get the first value from each pair and handle the error
	pi, err1 := p[i].GetFirst()
	pj, err2 := p[j].GetFirst()

	// If there's an error in either, we can handle it here (return false or other logic)
	if err1 != nil || err2 != nil {
		// Handling the error: you can choose to return false, true, or custom logic here
		return false
	}

	// Type assertion: pi and pj are of type `any`, we expect them to be of type `int` for comparison
	return pi.(int) < pj.(int) // Lower value = higher priority
}

// byB is a custom type to sort Pair items by field B in ascending order.
// A lower value in field B has a higher priority.
type byB []pair.Pair

func (p byB) Len() int      { return len(p) }
func (p byB) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byB) Less(i, j int) bool {
	// Get the second value from each pair and handle the error
	pi, err1 := p[i].GetSecond()
	pj, err2 := p[j].GetSecond()

	// If there's an error in either, we can handle it here (return false or other logic)
	if err1 != nil || err2 != nil {
		// Handling the error: you can choose to return false, true, or custom logic here
		return false
	}

	// Type assertion: pi and pj are of type `any`, we expect them to be of type `int` for comparison
	return pi.(int) < pj.(int) // Lower value = higher priority
}

// Enqueue adds a new Pair item to the queue.
func (q *PriorityQueue) Enqueue(item pair.Pair) {
	q.Data = append(q.Data, item)
}

// SortByA sorts the queue in ascending order based on field A (lower value = higher priority).
func (q *PriorityQueue) SortByA() {
	sort.Sort(byA(q.Data))
}

// SortByB sorts the queue in ascending order based on field B (lower value = higher priority).
func (q *PriorityQueue) SortByB() {
	sort.Sort(byB(q.Data))
}

// Dequeue removes and returns the first element in the queue (the one with the highest priority).
// It returns false if the queue is empty.
func (q *PriorityQueue) Dequeue() (pair.Pair, bool) {
	if len(q.Data) == 0 {
		return pair.Pair{}, false
	}
	item := q.Data[0]
	q.Data = q.Data[1:]
	return item, true
}

func (q *PriorityQueue) Peek() pair.Pair {
	if len(q.Data) == 0 {
		// Return the zero value of pair.Pair
		return pair.Pair{}
	}
	return q.Data[0]
}

func (q *PriorityQueue) Clear() {
	q.Data = nil
}

func (q *PriorityQueue) Length() int {
	return len(q.Data)
}
