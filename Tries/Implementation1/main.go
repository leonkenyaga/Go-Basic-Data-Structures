package main

import "fmt"

//AlphabetSIze is the numer of possible characters in the trie
const AlphabetSize = 26
//Node represents each node in the trie
type Node struct{
	children [AlphabetSize]*Node
	isEnd bool
}
//Trie
type Trie struct{
   root *Node
}
//Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength:=len(w)
	currentNode:=t.root
	for i:=0; i < wordLength; i++ {
		charIndex:= w[i]-'a'
      if currentNode.children[charIndex] == nil{
         currentNode.children[charIndex] = &Node{}
	  }
	  currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}
//Search will take in word and return true if that word is included in the Trie

func (t *Trie) Search(w string) bool {
	wordLength:=len(w)
	currentNode:=t.root
	for i:=0; i<wordLength; i++ {
		charIndex:= w[i]-'a'
      if currentNode.children[charIndex] == nil{
         return false
	  }
	  currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
      return true
	}
	return false
}


func InitTrie() *Trie {
	return &Trie{root:&Node{}}
}

func main(){
	myTrie:=InitTrie()
	fmt.Println(myTrie.root)
	toAdd:= []string{
		"aragorn",
		"aragog",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}
	for _, v:= range toAdd{
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.root)
	fmt.Println(myTrie.Search("eragon"))
}