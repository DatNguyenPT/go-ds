package pair

import "fmt"

type Pair struct {
	a, b any // No need for pointers
}

// NewPair creates and returns a new Pair with the given values.
func NewPair(a, b any) Pair {
	return Pair{a: a, b: b}
}

// GetFirst returns the first value if present, otherwise an error.
func (pair *Pair) GetFirst() (any, error) {
	if pair.a != nil {
		return pair.a, nil
	}
	return nil, fmt.Errorf("there is no first value")
}

// GetSecond returns the second value if present, otherwise an error.
func (pair *Pair) GetSecond() (any, error) {
	if pair.b != nil {
		return pair.b, nil
	}
	return nil, fmt.Errorf("there is no second value")
}

// RemoveFirst removes the first value if present.
func (pair *Pair) RemoveFirst() (bool, error) {
	if pair.a != nil {
		pair.a = nil
		return true, nil
	}
	return false, fmt.Errorf("no first element exists")
}

// RemoveSecond removes the second value if present.
func (pair *Pair) RemoveSecond() (bool, error) {
	if pair.b != nil {
		pair.b = nil
		return true, nil
	}
	return false, fmt.Errorf("no second element exists")
}
