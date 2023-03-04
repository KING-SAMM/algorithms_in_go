package main

import "fmt"

/*--------------------------------------------
| Connected Component Count
---------------------------------------------*/
/*
Write a function, connectedComponentsCount, that takes in the adjacency list of an undirected graph. The function should return the number of connected components within the graph.

* https://structy.net/problems/connected-comoponents-count
*/

func connectedComponentsCount(graph map[int][]int) int {
	visited := make(map[int]struct{})
	count := 0
	for node := range graph {
		if explore(graph, node, visited) == true {
			count += 1
		}
	}
	return count
}

func explore(graph map[int][]int, current int, visited map[int]struct{}) bool {
	if has(visited, current) { return false }
	visited[current] = struct{}{}

	for _, neighbor := range graph[current] {
		explore(graph, neighbor, visited)
	}
	return true
}

// Check if an element exists in a map 
func has(set map[int]struct{}, entry int) bool {
	for key, _ := range set {
		if (key == entry) {
			return true
		}
	}
	return false
}

func main() {
	// Build graph adjacency list
	graph := make(map[int][]int)
	a, b, c, d, e, f, g := 0, 1, 2, 3, 4, 5, 8
	
	graph[a] = append(graph[a], 8, 1, 5)
	graph[b] = append(graph[b], 0)
	graph[f] = append(graph[f], 0, 8)
	graph[g] = append(graph[g], 0, 5)
	graph[c] = append(graph[c], 3, 4)
	graph[d] = append(graph[d], 2, 4)
	graph[e] = append(graph[e], 3, 2)

	fmt.Println(graph)
	count := connectedComponentsCount(graph)
	fmt.Printf("The count is %v", count)
}