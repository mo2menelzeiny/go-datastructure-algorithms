package stacksandqueues

type LinkedNode struct {
	value int
	next  *LinkedNode
}

type Stack struct {
	top *LinkedNode
}

func (s *Stack) Push(value int) {
	node := LinkedNode{value: value}
	node.next = s.top
	s.top = &node
}

func (s *Stack) Pop() *LinkedNode {
	node := s.top
	s.top = s.top.next
	node.next = nil
	return node
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}
