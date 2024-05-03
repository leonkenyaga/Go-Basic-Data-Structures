package main

import "fmt"

type Node struct{
	Key   int
	Left  *Node
	Right *Node
}

var Count int 

func(n *Node) Insert(k int){
	if n.Key < k {
		if n.Right == nil {
			n.Right= &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key:k}
		} else {
        n.Left.Insert(k)
		}
	}
} 

func (n *Node) Search(k int) bool {
    Count++
	if n==nil{
		return false
	}
	if n.Key < k{
		return n.Right.Search(k)
	}else if n.Key >k {
        return n.Left.Search(k)
	}
	return true
}

func main(){
	tree:= &Node{ Key: 100 }
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(105)
	tree.Insert(165)
	tree.Insert(15)
	fmt.Println( *tree.Right, *tree.Right.Right,*tree.Right.Left, *&tree.Left, *tree.Left.Left)
	fmt.Println(tree.Search(165), Count)

}