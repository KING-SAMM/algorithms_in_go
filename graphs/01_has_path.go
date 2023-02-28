package main

import "fmt"

/*---------------------------------
|    HAS PATH
-----------------------------------*/
/**
 * PROBLEM STATEMENT:
 * 
 * Write a function, hasPath, that takes in an object representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the 'source' and 'destination' nodes.
 * 
 * https://structy.net/problems/has-path
 */


//1. Iterative breadth-first solution
// func hasPath(graph map[string][]string, src string, dst string) bool {
//     queue := []string{src}

//     for len(queue) > 0 {
//         current := queue[0]
//         queue = queue[1:]

// 		fmt.Printf("Queue is %v\n", queue)

//         if current == dst { return true }

//         for _, neighbor := range graph[current] {
//             queue = append(queue, neighbor)
//         }
//     }
//     return false
// }

//2. Recursive depth-first solution
// func hasPath(graph map[string][]string, src string, dst string) bool {
//     if src == dst { return true }

// 	for _, neighbor := range graph[src] {
// 		if hasPath(graph, neighbor, dst) == true {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	
	graph := make(map[string][]string)
	graph["f"] = []string{"g", "i"}
	graph["g"] = []string{"h"}
	graph["h"] = []string{}
	graph["i"] = []string{"g", "k"}
	graph["j"] = []string{"i"}
	graph["k"] = []string{}

	hasPath(graph, "f", "k")
}