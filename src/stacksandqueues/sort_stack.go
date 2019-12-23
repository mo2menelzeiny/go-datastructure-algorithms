package stacksandqueues

func SortStack(stack *Stack) {
	tempStack := Stack{}
	for !stack.isEmpty() {
		temp := stack.Pop()
		if !tempStack.isEmpty() && tempStack.top.value > temp.value {
			stack.Push(tempStack.Pop().value)
		}

		tempStack.Push(temp.value)
	}

	for !tempStack.isEmpty() {
		stack.Push(tempStack.Pop().value)
	}
}
