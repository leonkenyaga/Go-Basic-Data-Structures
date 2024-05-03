// Implementation of a HashTable (Basic) With Separate Chaining

package main

import (
	//"fmt"
	//"hash"
	"log"
)

const ArraySize = 7 //size of the HashTable array

//HashTable structure will hold an array

type HashTable struct{
	array [ArraySize] *bucket
}

// bucket is a linked list in each slot of the HashTable array

type bucket struct{
	head *bucketNode
}

//bucketNode is a linked list node that holds the key

type bucketNode struct{

	key string
	next *bucketNode

}



//Insert will take in a key and add it to the HashTable array
func (h *HashTable) Insert(key string){
	index:=hash(key)
	h.array[index].insert(key) // Inserts a key in a particular slot
}

//Search will take in a key and return true if that key is stored in the HashTable

func (h *HashTable) Search(key string) bool {
	index:=hash(key)
	return h.array[index].Search(key)
}

//Delete will take in a key and delete it from the hashtable

func (h *HashTable) Delete(key string){
	index:=hash(key)

	h.array[index].Delete(key)

}
//Hash Function
func hash(key string) int{
    sum:=0
	for _, v:=range key{
		
		sum+= int(v)

	}
	return sum % ArraySize //gets remainder

}

//Insert will take a key, create a node with the key and insert the node in the bucket 

func (b *bucket) insert(k string){
	//get a key and create a bucketNode with that key
	if !b.Search(k){
	newNode:= &bucketNode{key: k}
	newNode.next = b.head
	b.head=newNode
	} else {
		log.Println(k," already exists")
	}
}

// Search will take a key and return true if the bucket has the key 
func (b *bucket) Search(k string) bool {
    currentNode:= b.head
	//loop untill the currentNode is empty
	for currentNode != nil{
		if currentNode.key==k{
			return true
		}
		currentNode = currentNode.next//Routine Update
	}
	return false
}

// Delete will take in a key andf delete the node from the bucket
func (b *bucket) Delete(k string){
	//Linkedlist will be deleting the current nod e by skipping the currentnode and making the previous node point to the next node

	if b.head.key==k{ //if the key of the head is actually the match
		b.head = b.head.next
		return
	}

	PreviousNode:= b.head

	for PreviousNode.next!=nil{
		if PreviousNode.next.key==k{
			//delete
			PreviousNode.next=PreviousNode.next.next
		}
	}
}
//Init will create a bucket in each slot of the HashTable

func Init() *HashTable{
	result:=&HashTable{}

	for i:= range result.array{
		result.array[i]=&bucket{}
	}
	return result
}

func main(){
	TestHashTable:=Init()
	TestHashTable.Insert("Mark")
	TestHashTable.Insert("Logan")
	TestHashTable.Insert("Kendall")
	TestHashTable.Insert("Siobhan")
	TestHashTable.Insert("Roman")
	TestHashTable.Insert("Connor")
	log.Println(TestHashTable.Search("Connor"))
	log.Println(TestHashTable.Search("Logan"))
	log.Println(TestHashTable.Search("Siobhan"))
	TestHashTable.Delete("Connor")
	log.Println(TestHashTable.Search("Connor"))
	TestHashTable.Delete("Roman")
	log.Println(TestHashTable.Search("Kendall"))
	log.Println(TestHashTable.Search("Roman"))
}

