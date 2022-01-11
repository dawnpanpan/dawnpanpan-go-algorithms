package stack

import "fmt"


type Node struct {
	Value interface{}
	Next *Node
}

type Stack struct {
	top *Node
	size int
}

func (stack *Stack) Push(value interface{}) {
	newNode := &Node{} // create new node
	
	newNode.Value = value
	newNode.Next = stack.top

	stack.top = newNode
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	popValue := stack.top.Value
	if stack.top.Next == nil {
		stack.top = nil
	}else {
		stack.top.Value, stack.top.Next = stack.top.Next.Value, stack.top.Next.Next 
	}
	stack.size--
	return popValue
}

func (stack *Stack) Peek() interface{} {	
	return  stack.top.Value
}

func (stack *Stack) StackSize() interface{} {
	return stack.size
}

func (stack *Stack) IsEmpty() interface{} {
	return stack.StackSize() == 0
}

func (stack *Stack) Show() (in []interface{}) {
	current := stack.top

	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}

func  (stack *Stack) Test() {
	myStack := &Stack{}
	myStack.Push(4)
	myStack.Push(7)
	myStack.Push(5)
	myStack.Push(6)
	myStack.Pop()
	fmt.Println(myStack.Show())

}