package main

import "errors"

// filo is First-In-Last-Out structure aka "stack"
type filo struct {
	stack []rune
}

// Add character to the stack
func (f *filo) push(r rune) {
	f.stack = append(f.stack, r)
}

// pop removes character from the stack
func (f *filo) pop() (rune, error) {
	length := len(f.stack)
	if length == 0 {
		return 0, errors.New("Stack is empty")
	}
	Pop := f.stack[len(f.stack)-1]
	f.stack = f.stack[:len(f.stack)-1]
	return Pop, nil
}

// getLast returns the last character from the stack
func (f *filo) getLast() (rune, error) {
	length := len(f.stack)
	if length == 0 {
		return 0, errors.New("stack is empty")
	}
	return f.stack[length-1], nil
}
