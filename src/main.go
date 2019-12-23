package main

import "stacksandqueues"

func main() {
	TestStack()
	TestQueue()
	TestMultiStack()
	TestStackMin()
	TestStackOfStacks()
	TestTwoStacksQueue()
	TestSortStack()
}

func TestSortStack() {
	stack := stacksandqueues.Stack{}
	stack.Push(2)
	stack.Push(3)
	stack.Push(1)
	stack.Push(0)
	stacksandqueues.SortStack(&stack)
}

func TestTwoStacksQueue() {
	queue := stacksandqueues.TwoStacksQueue{}
	queue.Add(0)
	queue.Add(1)
	queue.Add(2)

	queue.Remove()
}

func TestStackOfStacks() {
	stack := stacksandqueues.StackOfStacks{Threshold: 3}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)

	stack.PopAt(2)
}

func TestStackMin() {
	stack := stacksandqueues.MinStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(0)
	stack.Min()
}

func TestMultiStack() {
	stack := stacksandqueues.MultiStack{}
	stack.Init(10, 4)
	stack.Push(2, 2)
	stack.Push(1, 5)
	stack.Push(3, 4)
	stack.Pop(1)
	stack.Push(2, 33)
	stack.Pop(2)
	stack.Push(4, 33)
}

func TestStack() {
	stack := stacksandqueues.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
}

func TestQueue() {
	queue := stacksandqueues.Queue{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Remove()
}
