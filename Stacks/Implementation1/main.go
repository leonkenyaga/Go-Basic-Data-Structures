package main

import (
	"fmt"
//	"log"
)

type Stack struct{
	items []int
}

//Push will add a value at the end

func (s *Stack) Push( i int){
	s.items = append(s.items, i)
}

// Pop will remove value at the end and returns the removed value

func (s *Stack) Pop() int {
	l:=len(s.items)-1
	toRemove:=s.items[l]
	s.items=s.items[:l]
	return toRemove
}

func main(){
	myStack:=Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack)
	for i:=0; i<4; i++{
		fmt.Println("poped out", myStack.Pop())
	}
	//myStack.Pop()
	fmt.Println(myStack)
}