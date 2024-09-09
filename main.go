package main

import (
	"fmt"
)

func main() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	fmt.Println(myStack.Peek())
	fmt.Println(myStack.Pop())

	fmt.Println(myStack.Peek())
	fmt.Println(myStack)

	myStack.Pop()
	myStack.Push(4)
	fmt.Println(myStack.Peek())
	fmt.Println(myStack)
}
