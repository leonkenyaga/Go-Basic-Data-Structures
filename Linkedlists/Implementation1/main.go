package main

import "fmt"

type node struct{
	data int   //value
	next *node //address of the next node on the list
}

type linkedlist struct{
	head *node //address of the first node in the linkedlist
	length int  //indicates how long the linked list is
}

//puts a node inside the linkedlist

func (l *linkedlist) prepend(n *node){
temporaryNode:=l.head //makes a temporary node holder and copies the pointer to the firstnode to it
	//n.next=l.head
	//l.head=n 
l.head = n //overrides the pointer to the firstnode with the pointer to a new node (the new first)
l.head.next = temporaryNode // tells the firstnode to point to the node it replaced (the old first)
    l.length ++ // increases the length of the linkedlist by one
}

// deletes a node with the given value from the linkedlist
func (l *linkedlist) deleteWithValue(value int){

	if l.length == 0{ //empty list case
     return
	}
	// if the input value turns out to be in the head data 
	if l.head.data == value{
		l.head = l.head.next // replace the pointer to the first node with a pointer the next node
		l.length--
		return
	}

	previousToDelete:=l.head
	
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil { // we'll assume tha we've reached the end and that there is no match for the input value
			return
		}
		previousToDelete = previousToDelete.next // hops to the next node 
	}
	// if the for loop condition ceases to be true  we can update the pointer so that it skips the node that needs to be deleted
	// the skipped node will be Garbage Collected
    previousToDelete.next = previousToDelete.next.next

	// Reduce the length of the linkedlist
     
	l. length--

}

func (l linkedlist) printListData() {
	toPrint:=l.head // initializes the first value to be printed as the firstnode
    //counts down until the length of trhe linkedlist  becomes zero
	for l.length !=0{
		fmt.Printf("%d ", toPrint.data)//prints the value on the current node
		toPrint=toPrint.next // updates the node to the next one on the list
		l.length-- //reduces the length of the list by one each time a value on the list has been printed
	}
	fmt.Println()
}

func main(){
	newList:=linkedlist{}
	node1:=&node{data: 10}
	node2:=&node{data: 7}
	node3:=&node{data:5}
	node4:=&node{data:6}
	node5:=&node{data:3}
	newList.prepend(node1)
	newList.prepend(node2)
	newList.prepend(node3)
	newList.prepend(node4)
	newList.prepend(node5)
	newList.printListData()
	newList.deleteWithValue(5)
	newList.printListData()
}