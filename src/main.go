package main

import "datastructure"

func main() {
	TestStack()
	TestQueue()
	TestMultiStack()
	TestStackMin()
}

func TestStackMin() {
	stack := datastructure.MinStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(0)
	stack.Min()
}

func TestMultiStack() {
	stack := datastructure.MultiStack{}
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
	stack := datastructure.Stack{}
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	stack.Pop()
}

func TestQueue() {
	queue := datastructure.Queue{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Remove()
}
