package main

import (
	"sync"
)

var lock sync.Mutex

type Stack struct {
	items []int
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() (bool, int) {
	if len(s.items) == 0 {
		return true, 0
	}
	return false, len(s.items)
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	lock.Lock()
	defer lock.Unlock()
	s.items = append(s.items, item)
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (int, bool) {
	lock.Lock()
	defer lock.Unlock()
	value, length := s.IsEmpty()
	if value {
		return 0, false
	}

	res := s.items[length-1]
	s.items = s.items[:length-1]
	return res, true
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (int, bool) {
	lock.Lock()
	defer lock.Unlock()
	value, length := s.IsEmpty()
	if value {
		return 0, false
	}

	return s.items[length-1], true

}
