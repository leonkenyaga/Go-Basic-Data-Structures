package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
//generates an array of random integers

func generateUniqueRandomNumbers(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	set := make(map[int]bool)
	var result []int
	for len(set) < n {
	   value := rand.Intn(max)
	   if !set[value] {
		  set[value] = true
		  result = append(result, value)
	   }
	}
	return result
 }

// CompareArray Compares the first element of a slice against all elements and returns a boolean value

func CompareArray( array []int ) bool {
	
	firstElement:= array[0]

	if len(array)== 1{
		return true
	}
	if length:= len(array); length > 1{

		for i:=1; i < length; i++{
		// returns false if the first element is not the largest
        if array[i] > firstElement{
			return false
		}
		}
		//returns true if the first element is the largest
		return true
	}
	return true
	
}
 

// TestHeapInsertion tests the accuracy of the Insert and Heapify Abstract Data Type operations
func TestHeapInsertion(t *testing.T) {
   
    array:=generateUniqueRandomNumbers(10, 100)
    
     heap:=&MaxHeap{}

	for _, v:=range array{
	 heap.Insert(v)
	 fmt.Println(heap)
	}


    if !CompareArray(heap.array){
        t.Error(heap, ` is not a heap  `)
    }
}
