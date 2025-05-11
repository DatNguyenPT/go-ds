package queue

type Queue struct {
	Data []any
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(element any) {
	q.Data = append(q.Data, element)
}

// Dequeue removes and returns the front element, or nil if empty
func (q *Queue) Dequeue() any {
	if len(q.Data) == 0 {
		return nil
	}
	front := q.Data[0]
	q.Data = q.Data[1:]
	return front
}

// Peek returns the front element without removing it
func (q *Queue) Peek() any {
	if len(q.Data) == 0 {
		return nil
	}
	return q.Data[0]
}

// Clear removes all elements
func (q *Queue) Clear() {
	q.Data = nil
}

// Length returns the number of elements in the queue
func (q *Queue) Length() int {
	return len(q.Data)
}
