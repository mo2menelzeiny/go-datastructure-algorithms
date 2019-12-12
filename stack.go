package main

type StackNode struct {
	value int
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

func (s Stack) add(value int) {
	node := StackNode{value: value}
	node.next = s.top
	s.top = &node
}
