package main

import "fmt"

type Graph struct {
	table map[int][]int
}
// type graph struct{
// 	starting *graphNode
// }

func main() {
	graph := &Graph{table: make(map[int][]int)}

	graph.insert(3, 5, true)
	graph.insert(3, 4, true)
	graph.insert(5, 6, false)
	graph.insert(4,5,true)

	 bfs := graph.BFS(3)
	// dfs := graph.DFS(3)
	fmt.Println(bfs, "\n ")

	graph.display()

}

func (g *Graph) display() {
	for vertex, edges := range g.table {
		fmt.Println(vertex, "  : ", edges)
	}
}

func (g *Graph) insert(vertex, edge int, bidirectional bool) {
	if !g.exist(vertex) {
		g.addVertex(vertex)
	}
	if !g.exist(edge) {
		g.addVertex(edge)
	}

	g.table[vertex] = append(g.table[vertex], edge)

	if bidirectional {
		g.table[edge] = append(g.table[edge], vertex)
	}

}

func (g *Graph) addVertex(vertex int) {
	// edges := new([]int)
	g.table[vertex] = []int{}
}

func (g *Graph) exist(vertex int) bool {
	_, exist := g.table[vertex]
	return exist
}

func (g *Graph) BFS(start int) []int {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true
	result := []int{}

	for len(queue) > 0 {

		vertex := queue[0]
		result = append(result, vertex)
		queue = queue[1:]
		for _, v := range g.table[vertex] {

			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	for key := range g.table {

		if !visited[key] {
			g.BFS(key)
		}
	}
	return result
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{start}

	var DFSUTIL func(vertex int)
	DFSUTIL = func(vertex int) {

		visited[vertex] = true``
		for _, v := range g.table[vertex] {

			if !visited[v] {

				DFSUTIL(v)
				result = append(result, v)
			}
		}

	}

	DFSUTIL(start)
	for key := range g.table {
		if !visited[key] {
			g.DFS(key)
		}
	}
	return result
}
