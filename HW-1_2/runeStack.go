package main

import "errors"

//RuneStack ...
type RuneStack struct {
	data []rune
}

//NewRuneStack ...
func NewRuneStack() *RuneStack {
	return &RuneStack{make([]rune, 0)}
}

//Push ...
func (s *RuneStack) Push(v rune) {
	s.data = append(s.data, v)
}

//Pop ...
func (s *RuneStack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("Incorrect expression")
	}
	l := len(s.data)
	res := s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

//Top ...
func (s *RuneStack) Top() rune {
	return s.data[len(s.data)-1]
}

//IsEmpty ...
func (s *RuneStack) IsEmpty() bool {
	return len(s.data) == 0
}
