// Implementation of a Graph (Basic) With an Adjacency List

package main

import "fmt"

//Graph represents an adjacency list graph

type Graph struct{
	
	vertices []*Vertex

}

//Vertex represents a Graph vertex

type Vertex struct{

	key int
	adjacent []*Vertex
}

//AddVertex adds a Vertex to the Graph

func (g *Graph) AddVertex(k int){
	if contains(g.vertices, k){
		err:=fmt.Errorf("vertex %v not added because it is an existing key", k)
        fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//AddEdge adds an edge to the graph

func (g *Graph) AddEdge(from, to int){
	
	//get Vertex
	fromVertex:=g.getVertex(from)
	toVertex:=g.getVertex(to)

    // Check error
	if fromVertex==nil || toVertex == nil{
		err:=fmt.Errorf("Invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent, to){
        err:= fmt.Errorf("Existing edge (%v --> %v)",from, to)
		fmt.Println(err.Error())
	}else{
		// add Edge
		fromVertex.adjacent=append(fromVertex.adjacent, toVertex)
	}
}

//getVertex returns a pointer to the Vertex with a key integer

func (g *Graph) getVertex(k int) *Vertex {
	for i,v:= range g.vertices{
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//Contains

func contains( s []*Vertex, k int) bool {
	for _, v := range s{
		if k ==v.key{
			return true
		}
	}
	return false
}

// Print will print the adjacent list from each Vertex of the Graph

func (g *Graph) print(){

	for _,V:= range g.vertices{
		fmt.Printf("\n Vertex  %v: ", V.key)

		for _, K:= range V.adjacent{
			fmt.Printf("%v ", K.key)
		}
	}

    fmt.Println()

}

func main(){

	Graph:= &Graph{}
	for i := 0; i < 5; i++ {
		Graph.AddVertex(i)
	}
	Graph.print()
	Graph.AddEdge(3,1)
	Graph.AddEdge(3,1)
	Graph.AddEdge(4,3)
	Graph.AddEdge(2,2)
	Graph.print()

}