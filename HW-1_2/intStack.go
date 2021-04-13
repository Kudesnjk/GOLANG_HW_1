package main

import "errors"

//IntStack ...
type IntStack struct {
	data []int
}

//NewIntStack ...
func NewIntStack() *IntStack {
	return &IntStack{make([]int, 0)}
}

//Push ...
func (s *IntStack) Push(v int) {
	s.data = append(s.data, v)
}

//Pop ...
func (s *IntStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Incorrect expression")
	}
	l := len(s.data)
	res := s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

//Top ...
func (s *IntStack) Top() int {
	return s.data[len(s.data)-1]
}

//IsEmpty ...
func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}
