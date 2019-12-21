package datastructure

/* Stack Min: How would you design a stack which, in addition to push and pop, has a function min
 * which returns the minimum element? Push, pop and min should all operate in 0(1) time.
 */

type MinLinkedNode struct {
	min   int
	value int
	next  *MinLinkedNode
}

type MinStack struct {
	top *MinLinkedNode
}

func (m *MinStack) Push(value int) {
	node := MinLinkedNode{min: value, value: value, next: m.top}
	if m.top != nil && m.top.min < node.min {
		node.min = m.top.min
	}

	m.top = &node
}

func (m *MinStack) Pop() *MinLinkedNode {
	node := m.top
	m.top = m.top.next
	node.next = nil
	return node
}

func (m *MinStack) Min() int {
	if m.top != nil {
		return m.top.min
	}

	return 0
}
