package main

import "fmt"

/*--------------------------------------------
| Largest Component
---------------------------------------------*/
/**
 * Write a function, largestComponent, that takes in the adjacency list of 
 * an undirected graph. The function should return the size of the largest
 * connected component in the graph.
 * 
 * @param {obj} graph 
 * @returns longest
 * 
 * https://structy.net/problems/largest-component
 */

 func largestComponent(graph map[string][]string) int {
	visited := make(map[string]struct{})
	largest := 0

	for node, _ := range graph {
		size := exploreSize(graph, node, visited)
		if largest < size {
			largest = size
		}
	}
	return largest
 }

 func exploreSize(graph map[string][]string, current string, visited map[string]struct{}) int {
	if has(visited, current) { return 0 }
	visited[current] = struct{}{}

	size := 1
	for _, neighbor := range graph[current] {
		size += exploreSize(graph, neighbor, visited)
	}
	return size
 }

// Check if an element exists in a set 
func has(set map[string]struct{}, entry string) bool {
	for key, _ := range set {
		if (key == entry) {
			return true
		}
	}
	return false
}

func main() {
	// Build graph adjacency list
	graph := make(map[string][]string)
	a, b, c, d, e, f := "1", "2", "6", "9", "7", "8"
	
	graph[a] = append(graph[a], "2")
	graph[b] = append(graph[b], "1", "8")
	graph[c] = append(graph[c], "7")
	graph[d] = append(graph[d], "8")
	graph[e] = append(graph[e], "6", "8")
	graph[f] = append(graph[f], "9", "7", "2")

	fmt.Println(graph)
	count := largestComponent(graph)
	fmt.Printf("The count is %v", count)

	// largestComponent({
	// 	0: ['8', '1', '5'],
	// 	1: ['0'],
	// 	5: ['0', '8'],
	// 	8: ['0', '5'],
	// 	2: ['3', '4'],
	// 	3: ['2', '4'],
	// 	4: ['3', '2']
	//   }); // -> 4

	// largestComponent({
	// 	1: ['2'],
	// 	2: ['1','8'],
	// 	6: ['7'],
	// 	9: ['8'],
	// 	7: ['6', '8'],
	// 	8: ['9', '7', '2']
	//   }); // -> 6

	// largestComponent({
	// 	3: [],
	// 	4: ['6'],
	// 	6: ['4', '5', '7', '8'],
	// 	8: ['6'],
	// 	7: ['6'],
	// 	5: ['6'],
	// 	1: ['2'],
	// 	2: ['1']
	//   }); // -> 5

	// largestComponent({}); // -> 0

	// largestComponent({
	// 	0: ['4','7'],
	// 	1: [],
	// 	2: [],
	// 	3: ['6'],
	// 	4: ['0'],
	// 	6: ['3'],
	// 	7: ['0'],
	// 	8: []
	//   }); // -> 3
}
