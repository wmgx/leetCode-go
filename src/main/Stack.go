package main

import "errors"

type Item interface{}

type Stack struct {
	size     int
	elements []Item
	top      int
}

func NewStackWithSize(size int) *Stack {
	return &Stack{size, make([]Item, size), 0}
}
func NewStack() *Stack {
	return NewStackWithSize(10)
}
func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}
func (stack *Stack) Push(e Item) {
	if stack.top == stack.size {
		var newArray = make([]Item, stack.size*2)
		copy(stack.elements, newArray)
		stack.elements = newArray
		stack.size *= 2
	}
	stack.elements[stack.top] = e
	stack.top++
}
func (stack *Stack) Pop() (Item, error) {
	if stack.IsEmpty() {
		return nil, errors.New("The stack is empty")
	}
	stack.top--
	return stack.elements[stack.top+1], nil
}
