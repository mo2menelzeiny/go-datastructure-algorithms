package stacksandqueues

type TwoStacksQueue struct {
	stackA Stack
	stackB Stack
	size   int
}

func (tsq *TwoStacksQueue) move() {
	if !tsq.stackB.isEmpty() {
		return
	}

	for tsq.stackA.isEmpty() == false {
		tsq.stackB.Push(tsq.stackA.Pop().value)
	}
}

func (tsq *TwoStacksQueue) Add(value int) {
	tsq.stackA.Push(value)
	tsq.size++
}

func (tsq *TwoStacksQueue) Remove() *LinkedNode {
	tsq.move()
	tsq.size--
	return tsq.stackB.Pop()
}
