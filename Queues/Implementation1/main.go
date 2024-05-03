package main

import "fmt"

//Queue reprents a queue that holds a slice

type Queue struct{
	items []int
}
// Enqueue adds a value at the end

func (q *Queue) Enqueue (i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the front and returns the removed value

func (q *Queue) Dequeue () int {
	toRemove:=q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main(){
	myQueue:=Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue(), myQueue)
	fmt.Println(myQueue.Dequeue(), myQueue)
	fmt.Println(myQueue.Dequeue(), myQueue)
	myQueue.Enqueue(4)
	fmt.Println(myQueue)
	myQueue.Enqueue(5)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue(), myQueue)
	fmt.Println(myQueue.Dequeue(), myQueue)
}