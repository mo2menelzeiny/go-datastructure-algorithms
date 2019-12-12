package main

import "datastructure"

func main() {
	TestStack()
	TestQueue()
}

func TestStack() {
	stack := datastructure.Stack{}
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	x := stack.Pop()
	println(x)
}

func TestQueue() {
	queue := datastructure.Queue{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	x := queue.Remove()
	println(x)
}
