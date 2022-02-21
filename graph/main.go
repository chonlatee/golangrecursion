package main

import (
	"container/list"
	"fmt"
)

// undirected graph

type graph struct {
	vertices map[string][]string
}

func (g *graph) Add(src, dst string) {

	if !contain(g.vertices[src], dst) {
		g.vertices[src] = append(g.vertices[src], dst)
	}

	if !contain(g.vertices[dst], src) {
		g.vertices[dst] = append(g.vertices[dst], src)
	}

}

func (g *graph) DFSTraversalRecursion(key string, visited map[string]bool) {
	fmt.Printf("start with vertice %v and have edge vertices %v\n", key, g.vertices[key])
	for _, v := range g.vertices[key] {
		if !visited[v] {
			visited[v] = true
			fmt.Printf("visit vertice %v\n", v)
			g.DFSTraversalRecursion(v, visited)
		}
	}
}

func (g *graph) DFSTraversal(key string) {

	st := list.New()
	visited := make(map[string]bool)

	st.PushBack(key)

	for st.Len() > 0 {
		e := st.Back()
		k := e.Value.(string)
		st.Remove(e)
		fmt.Printf("start with vertice %v and have edge vertices %v\n", k, g.vertices[k])
		for _, v := range g.vertices[k] {
			if !visited[v] {
				fmt.Printf("visit vertice %v\n", v)
				visited[v] = true
				st.PushBack(v)
				break
			}
		}
	}

}

func (g *graph) BFSTraversal(key string) {
	st := list.New()
	visited := make(map[string]bool)

	st.PushBack(key)

	visited[key] = true

	for st.Len() > 0 {
		e := st.Front()
		k := e.Value.(string)
		st.Remove(e)
		fmt.Printf("start with vertice %v and have edge vertices %v\n", k, g.vertices[k])
		for _, v := range g.vertices[k] {
			if !visited[v] {
				visited[v] = true
				fmt.Printf("visit vertice %v\n", v)
				st.PushBack(v)
			}
		}
	}
}

func contain(src []string, dst string) bool {
	for _, v := range src {
		if v == dst {
			return true
		}
	}

	return false
}

func NewGraph() *graph {
	return &graph{
		vertices: make(map[string][]string),
	}
}

func main() {

	g := NewGraph()

	g.Add("a", "b")
	g.Add("a", "c")
	g.Add("b", "c")
	g.Add("c", "d")
	g.Add("d", "e")
	g.Add("e", "c")
	g.Add("e", "f")

	startVertice := "f"

	visited := make(map[string]bool)
	visited[startVertice] = true

	g.DFSTraversalRecursion(startVertice, visited)
	fmt.Println()

	g.DFSTraversal(startVertice)
	fmt.Println()

	g.BFSTraversal(startVertice)
	fmt.Println()

}
