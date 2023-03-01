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

	return hasPath(graph, nodeA, nodeB, make(map[string]struct{}))
}

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

func has(set map[string]struct{}, entry string) bool {
	for key, _ := range set {
		if (key == entry) {
			return true
		}
	}
	return false
}


const buildEdges = (edges) => [
	const graph = {};

	for (let edge of edges) {
		const [a, b] = edge;

		if (!(a in graph)) graph[a] = [];
		if (!(b in graph)) graph[b] = [];

		graph[a].push(b);
		graph[b].push(a);
	}

	return graph;
]

