package datastructure

type StackOfStacksNode struct {
	next     *StackOfStacksNode
	subStack *Stack
	size     int
}

type StackOfStacks struct {
	top       *StackOfStacksNode
	Threshold int
}

func (ms *StackOfStacks) Push(value int) {
	// allocate new stack if size meets Threshold
	if ms.top == nil || ms.top.size == ms.Threshold {
		stackOfStacksNode := StackOfStacksNode{
			next:     ms.top,
			subStack: &Stack{top: nil},
			size:     0,
		}

		ms.top = &stackOfStacksNode
	}

	// push value to sub stack
	ms.top.subStack.Push(value)
	ms.top.size++

}

func (ms *StackOfStacks) Pop() *LinkedNode {
	// gate against empty stack
	if ms.top == nil {
		return nil
	}

	node := ms.top.subStack.Pop()
	ms.top.size--

	// deallocate if empty
	if ms.top.size == 0 {
		stackOfStacksNode := ms.top
		ms.top = ms.top.next
		stackOfStacksNode.next = nil
	}

	return node
}

func (ms *StackOfStacks) getStack(stackNo int) *StackOfStacksNode {
	// gate against invalid stack number
	if stackNo < 1 {
		return nil
	}

	node := ms.top
	for i := 0; i < stackNo-1; i++ {
		node = node.next
	}

	return node
}

func (ms *StackOfStacks) PopAt(stackNo int) *LinkedNode {
	// pop node from selected stack
	currSSNode := ms.getStack(stackNo)
	subStackNode := currSSNode.subStack.Pop()
	currSSNode.size--

	// rollover
	for i := stackNo - 1; i > 0; i-- {
		prevSSNode := ms.getStack(i)
		beforeTail := prevSSNode.subStack.top

		// reach item before bottom of prev stack
		// stack cannot be less than size - 1 because of rollover so no need for checking bounds
		for i := prevSSNode.size; i > 2; i-- {
			beforeTail = beforeTail.next
		}

		// remove tail node from prev stack
		tail := beforeTail.next
		beforeTail.next = nil
		prevSSNode.size--

		// rollover node to next stack
		prevSSNode.next.subStack.Push(tail.value)
		prevSSNode.next.size++
	}

	return subStackNode
}
