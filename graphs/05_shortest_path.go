package main

import "fmt"

/*--------------------------------------------
| Shortest Path
---------------------------------------------*/
/**
 * Write a function, shortestPath, that takes in an array of edges for an
 * undirected graph and two nodes (nodeA, nodeB). The function should return the
 * length of the shortest path between A and B. Consider the length as the number
 * of edges in the path, not the number of nodes. If there is no path between A
 * and B, then return -1.
 * 
 * @param {obj} graph 
 * @param {String} nodeA 
 * @param {String} nodeB 
 */


 func shortestPath(edges [][]string, nodeA string, nodeB string) int {
	graph := buildGraph(edges)
	visited := make(map[string]bool)
	queue := [][]interface{}{ {nodeA, 0} }	
	visited[nodeA] = true

	for len(queue) > 0 {
		node, distance := queue[0][0], queue[0][1]
		queue = queue[1:]

		if node == nodeB {
			return distance.(int)
		}

		for _, neighbor := range graph[node.(string)] {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				queue = append(queue, []interface{}{neighbor, distance.(int) + 1})
			}
		}
	}
	return -1
 }

 // Build up the graph of nodes from the edges
func buildGraph(edges [][]string) map[string][]string {
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
	/*-------------
	| EDGES A
	---------------*/
	edges := [][]string{ {"w", "x"}, {"x", "y"}, {"z", "y"}, {"z", "v"}, {"w", "v"} }

	// 	const edges = [
	//   ['w', 'x'],
	//   ['x', 'y'],
	//   ['z', 'y'],
	//   ['z', 'v'],
	//   ['w', 'v']
	// ];

	fmt.Println(shortestPath(edges, "w", "z")); // -> 2
	fmt.Println(shortestPath(edges, "y", "x")); // -> 1

	/*-------------
	| EDGES B
	---------------*/
	// edges := [][]string{ {"a", "c"}, {"a", "b"}, {"c", "b"}, {"c", "d"}, {"b", "d"}, {"e", "d"}, {"g", "f"} }

	// const edges = [
	//   ['a', 'c'],
	//   ['a', 'b'],
	//   ['c', 'b'],
	//   ['c', 'd'],
	//   ['b', 'd'],
	//   ['e', 'd'],
	//   ['g', 'f']
	// ];

	// fmt.Println(shortestPath(edges, "a", "e")); // -> 3
	// fmt.Println(shortestPath(edges, "e", "c")); // -> 2
	// fmt.Println(shortestPath(edges, "b", "g")); // -> -1


	/*-------------
	| EDGES C
	---------------*/
	// edges := [][]string{ {"c", "n"}, {"c", "e"}, {"c", "s"}, {"c", "w"}, {"w", "e"} }

	// const edges = [
	//   ['c', 'n'],
	//   ['c', 'e'],
	//   ['c', 's'],
	//   ['c', 'w'],
	//   ['w', 'e'],
	// ];

	// fmt.Println(shortestPath(edges, "w", "e")); // -> 1
	// fmt.Println(shortestPath(edges, "n", "e")); // -> 2


	/*-------------
	| EDGES D
	---------------*/
	// edges := [][]string{ {"m", "n"}, {"n", "o"}, {"o", "p"}, {"p", "q"}, {"t", "o"}, {"r", "q"}, {"r", "s"} }

	// const edges = [
	//   ['m', 'n'],
	//   ['n', 'o'],
	//   ['o', 'p'],
	//   ['p', 'q'],
	//   ['t', 'o'],
	//   ['r', 'q'],
	//   ['r', 's']
	// ];

	// fmt.Println(shortestPath(edges, "m", "s")); // -> 6

 }