package main

import "fmt"

/*--------------------------------------------
| Undirected Path
---------------------------------------------*/
/*
Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

* https://structy.net/problems/undirected-path
*/

func undirectedPath(edges []string, nodeA string, nodeB string) bool {
	graph := buildEdges(edges)

	return hasPath(graph, nodeA, nodeB, visited)
}

func hasPath(graph map[string][]string, src string, dst string, visited []string) bool {
    if src == dst { return true }
	if ()

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst) == true {
			return true
		}
	}
	return false
}

