package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	adjacency map[string][]string // every node is a key, and value is the list of its child node
}

func main() {
	g := NewGraph()
	g.AddVertex("harry")
	g.AddVertex("minhnam")
	g.AddVertex("boiboi")
	g.AddVertex("duong")

	g.AddEdge("harry", "minhnam")
	g.AddEdge("harry", "boiboi")
	g.AddEdge("minhnam", "duong")

	g.BFS("harry")
	g.ReDFS("harry")
	g.CreatePath("harry", "duong")
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, error := g.adjacency[vertex]; error {
		fmt.Printf("vertex already exist: %v", vertex)
		return false
	}

	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, error := g.adjacency[vertex]; !error {
		fmt.Printf("vertex is not exist: %v", vertex)
		return false
	}

	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node is already exist: %v", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)

	return true
}

func (g *Graph) BFS(startingNode string) {
	visited := g.initVisited()

	var queue []string
	visited[startingNode] = true
	queue = append(queue, startingNode)

	for len(queue) > 0 {
		var current string
		current, queue = queue[0], queue[1:] // dequeue, current is the first item of the queue
		fmt.Println("BFS", current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				queue = append(queue, node) // queue
				visited[node] = true
			}
		}
	}
}

func (g *Graph) CreatePath(firstNode, secondNode string) bool {
	visited := g.initVisited()
	var (
		path []string
		q    []string
	)
	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]

		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}

	fmt.Println("no link found")
	return false
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (g *Graph) ReDFS(startingNode string) {
	visited := g.initVisited()
	g.recursiveDFS(startingNode, visited)
}

func (g *Graph) DFS(startingNode string) {
	visited := g.initVisited()
	var stack []string
	visited[startingNode] = true
	stack = append(stack, startingNode)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// for _, node := range g.adjacency[current] {
		// 	if !visited[node] {
		// 		visited[node] = true
		// 		stack = append(stack, node)
		// 	}
		// }
	}
}

func (g *Graph) recursiveDFS(startingNode string, visited map[string]bool) {
	visited[startingNode] = true
	fmt.Println("DFS", startingNode)
	for _, node := range g.adjacency[startingNode] {
		if !visited[node] {
			visited[node] = true
			g.recursiveDFS(node, visited) // now the Vertex will be checked in depth
		}
	}
}

func (g *Graph) initVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}

	return visited
}
