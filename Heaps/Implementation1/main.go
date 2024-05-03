package main

import "fmt"

//MaxHeap struct has a slice that holds the array

type MaxHeap struct{
	array []int
}

//Insert adds an element to the heap
 func (h *MaxHeap) Insert (key int){
//Adding the key parameter into the array
	h.array=append(h.array,key)
//Pass the index of the last(added) element into the maxHeapifyUp method
    h.maxHeapifyUp(len(h.array)-1)
 }

//Extracts and returns the largest key and removes it from the heap
 func (h *MaxHeap) Extract ()int {
	if len(h.array)==0{
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	
    extracted:=h.array[0] //takes the first(largest) element
	//l stores the last index
	
	l:=len(h.array) - 1

	//when the array is empty (to avoid panic)
	
	//take out the last index and put it in the root
	h.array[0]=h.array[l]
	h.array= h.array[:l]

	h.maxHeapifyDown(0)

	return extracted

 }

 //maxHeapify method utilizes a bunch of helping functions for its operations
func (h *MaxHeap) maxHeapifyUp(index int){
	//When the parent of the element in question is smaller than its child(the element)
	for h.array[parent(index)] < h.array[index] {
	//SWap the parent's index (location) with the child's location (current index)
	h.Swap(parent(index), index)
	//Update the loop on the current index (location) of the element in question
	index= parent(index)
	} 
}

//maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int){
	lastIndex:=len(h.array)-1
	//get the indexes of its left and right children
	l,r :=left(index), right(index)
	childToCompare:=0
    //fmt.Println("This line")
	//loop while index has atleast one child (when the value of the left child index is within the range of the slice's elements)
    for l<=lastIndex{
		//fmt.Println("This line2")
		if l==lastIndex{//when the left child is the only child(if its the last element in the slice)
			childToCompare=l
		}else if h.array[l] > h.array[r]{//when the left child is larger
           childToCompare=l
		}else{//when right child is larger
        childToCompare=r
		}
	
	
	//Compare array value of current index to larger child and swap if smaller
	if h.array[index]<h.array[childToCompare]{
		//Swap with larger child
		h.Swap(index,childToCompare)
		//Update the index of the element in question and that of its children
		index=childToCompare
		l,r=left(index), right(index)
	}else{//stop the whole heapify process
		return

	}
}
}

//returns the parent's index

func parent(i int) int {
	
	return (i-1)/2
}
//return the left child's index

func left(i int) int{
	return 2*i +1
}

//returns the right child's index

func right(i int) int{
return 2*i + 2
}


//Swaps keys in the array (the parent and child element)
func(h *MaxHeap) Swap(i1, i2 int){

	h.array[i1],h.array[i2] = h.array[i2], h.array[i1]
}

func main(){
	array:=[]int{10,20,30,40,50, 60 , 70}
	//array=append(array, 43)
	fmt.Println(array)
	heap:=&MaxHeap{}
	//x:=heap.array[0]
	//fmt.Println(x)
	fmt.Println(heap)
	for _, v:=range array{
		heap.Insert(v)
	}
	
	fmt.Println("heap after insertion",heap)
	for i := 0; i < 20; i++ {
		fmt.Println("element extracted",heap.Extract())
		fmt.Println("element after extraction", heap)
	}
    
	fmt.Println(heap)
	
}
