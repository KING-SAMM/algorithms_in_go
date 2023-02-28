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
	visited = append(visited, src)

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst) == true {
			return true
		}
	}
	return false
}

const hasPath = (graph, src, dst, visited) => {
	if (src === dst) return true;
	if (visited.has(src)) return false;

	visited.add(src);

	for (let neighbor of graph[src]) {
		if (hasPath(graph, neighbor, dst, visited) === true) {
			return true;
		};
	};

	return false;
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

