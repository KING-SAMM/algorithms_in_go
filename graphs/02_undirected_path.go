package main

import "fmt"

/*--------------------------------------------
| Undirected Path
---------------------------------------------*/
/*
Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

* https://structy.net/problems/undirected-path
*/

func undirectedPath(edges [][]string, nodeA string, nodeB string) bool {
	graph := buildNodes(edges)

	// fmt.Println(hasPath(graph, nodeA, nodeB, make(map[string]struct{})))
	return hasPath(graph, nodeA, nodeB, make(map[string]struct{}))
}

// Implement hasPath algorithm
// Use map data structure (as a set) to keep track of visited nodes
// in order to avoid infinite loops from cyclic paths
func hasPath(graph map[string][]string, src string, dst string, visited map[string]struct{}) bool {
    if src == dst { return true }
	if has(visited, src) { return false }
	visited[src] = struct{}{}

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst, visited) == true {
			return true
		}
	}
	return false
}

// Check if an element exists in a map 
func has(set map[string]struct{}, entry string) bool {
	for key, _ := range set {
		if (key == entry) {
			return true
		}
	}
	return false
}

// Build up the graph of nodes from the edges
func buildNodes(edges [][]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		a, b := edge[0], edge[1]

		if _, ok := graph[a]; !ok {
			graph[a] = make([]string, 0)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make([]string, 0)
		}

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}

func main() {
	edges := make([][]string, 0)
	edge := []string{"i", "j"}
	edges = append(edges, edge)
	edge = []string{"k", "i"}
	edges = append(edges, edge)
	edge = []string{"m", "k"}
	edges = append(edges, edge)
	edge = []string{"k", "l"}
	edges = append(edges, edge)
	edge = []string{"o", "n"}
	edges = append(edges, edge)

	undirectedPath(edges, "j", "m")
}