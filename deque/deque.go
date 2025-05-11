package deque

type Deque struct {
	Data []any
}

// PushFront adds an element to the front
func (d *Deque) PushFront(element any) {
	d.Data = append([]any{element}, d.Data...)
}

// PushBack adds an element to the back
func (d *Deque) PushBack(element any) {
	d.Data = append(d.Data, element)
}

// PopFront removes and returns the front element
func (d *Deque) PopFront() any {
	if len(d.Data) == 0 {
		return nil
	}
	front := d.Data[0]
	d.Data = d.Data[1:]
	return front
}

// PopBack removes and returns the last element
func (d *Deque) PopBack() any {
	if len(d.Data) == 0 {
		return nil
	}
	back := d.Data[len(d.Data)-1]
	d.Data = d.Data[:len(d.Data)-1]
	return back
}

// PeekFront returns the front element without removing
func (d *Deque) PeekFront() any {
	if len(d.Data) == 0 {
		return nil
	}
	return d.Data[0]
}

// PeekBack returns the last element without removing
func (d *Deque) PeekBack() any {
	if len(d.Data) == 0 {
		return nil
	}
	return d.Data[len(d.Data)-1]
}

// Clear removes all elements
func (d *Deque) Clear() {
	d.Data = nil
}

// Length returns the number of elements
func (d *Deque) Length() int {
	return len(d.Data)
}
