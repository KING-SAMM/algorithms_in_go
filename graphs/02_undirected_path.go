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
	return hasPath(graph, nodeA, nodeB, make(map[string]bool))
}

// Implement hasPath algorithm
// Use map data structure (as a set) to keep track of visited nodes
// in order to avoid infinite loops from cyclic paths
func hasPath(graph map[string][]string, src string, dst string, visited map[string]bool) bool {
    if src == dst { return true }
	if has(visited, src) { return false }
	visited[src] = true

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst, visited) == true {
			return true
		}
	}
	return false
}

// Check if an element exists in a map 
func has(set map[string]bool, entry string) bool {
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

	/*-------------
	| EDGES A
	---------------*/
	edges = append(edges, []string{"i", "j"})
	edges = append(edges, []string{"k", "i"})
	edges = append(edges, []string{"m", "k"})
	edges = append(edges, []string{"k", "l"})
	edges = append(edges, []string{"o", "n"})

	fmt.Println(undirectedPath(edges, "j", "m")) // true
	// fmt.Println(undirectedPath(edges, "m", "j")) // true
	// fmt.Println(undirectedPath(edges, "l", "j")) // true
	// fmt.Println(undirectedPath(edges, "k", "o")) // false
	// fmt.Println(undirectedPath(edges, "i", "o")) // false

	// test_00:
	// 	const edges = [
	//   ['i', 'j'],
	//   ['k', 'i'],
	//   ['m', 'k'],
	//   ['k', 'l'],
	//   ['o', 'n']
	// ];

	// undirectedPath(edges, 'j', 'm'); // -> true

	// test_01:
	// const edges = [
	//   ['i', 'j'],
	//   ['k', 'i'],
	//   ['m', 'k'],
	//   ['k', 'l'],
	//   ['o', 'n']
	// ];

	// undirectedPath(edges, 'm', 'j'); // -> true
	// test_02:
	// const edges = [
	//   ['i', 'j'],
	//   ['k', 'i'],
	//   ['m', 'k'],
	//   ['k', 'l'],
	//   ['o', 'n']
	// ];

	// undirectedPath(edges, 'l', 'j'); // -> true
	// test_03:
	// const edges = [
	//   ['i', 'j'],
	//   ['k', 'i'],
	//   ['m', 'k'],
	//   ['k', 'l'],
	//   ['o', 'n']
	// ];

	// undirectedPath(edges, 'k', 'o'); // -> false
	// test_04:
	// const edges = [
	//   ['i', 'j'],
	//   ['k', 'i'],
	//   ['m', 'k'],
	//   ['k', 'l'],
	//   ['o', 'n']
	// ];

	// undirectedPath(edges, 'i', 'o'); // -> false

	/*-------------
	| EDGES B
	---------------*/
	// edges = append(edges, []string{"a", "b"})
	// edges = append(edges, []string{"c", "a"})
	// edges = append(edges, []string{"b", "c"})
	// edges = append(edges, []string{"q", "r"})
	// edges = append(edges, []string{"q", "s"})
	// edges = append(edges, []string{"q", "u"})
	// edges = append(edges, []string{"q", "t"})

	// fmt.Println(undirectedPath(edges, "a", "b")) // true
	// fmt.Println(undirectedPath(edges, "a", "c")) // true
	// fmt.Println(undirectedPath(edges, "r", "t")) // true
	// fmt.Println(undirectedPath(edges, "r", "b")) // false

	// test_05:
	// const edges = [
	//   ['b', 'a'],
	//   ['c', 'a'],
	//   ['b', 'c'],
	//   ['q', 'r'],
	//   ['q', 's'],
	//   ['q', 'u'],
	//   ['q', 't'],
	// ];


	// undirectedPath(edges, 'a', 'b'); // -> true
	// test_06:
	// const edges = [
	//   ['b', 'a'],
	//   ['c', 'a'],
	//   ['b', 'c'],
	//   ['q', 'r'],
	//   ['q', 's'],
	//   ['q', 'u'],
	//   ['q', 't'],
	// ];

	// undirectedPath(edges, 'a', 'c'); // -> true
	// test_07:
	// const edges = [
	//   ['b', 'a'],
	//   ['c', 'a'],
	//   ['b', 'c'],
	//   ['q', 'r'],
	//   ['q', 's'],
	//   ['q', 'u'],
	//   ['q', 't'],
	// ];

	// undirectedPath(edges, 'r', 't'); // -> true
	// test_08:
	// const edges = [
	//   ['b', 'a'],
	//   ['c', 'a'],
	//   ['b', 'c'],
	//   ['q', 'r'],
	//   ['q', 's'],
	//   ['q', 'u'],
	//   ['q', 't'],
	// ];

	// undirectedPath(edges, 'r', 'b'); // -> false

	/*-------------
	| EDGES C
	---------------*/
	// edges = append(edges, []string{"s", "r"})
	// edges = append(edges, []string{"t", "q"})
	// edges = append(edges, []string{"q", "r"})

	// fmt.Println(undirectedPath(edges, "r", "t")) // true
		
	// test_09:
	// const edges = [
	//   ['s', 'r'],
	//   ['t', 'q'],
	//   ['q', 'r'],
	// ];

// undirectedPath(edges, 'r', 't'); // -> true
}