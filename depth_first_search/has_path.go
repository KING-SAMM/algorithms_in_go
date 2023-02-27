package main

import "fmt"

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

func hasPath(graph map[string][]string, src string, dst string) bool {
    if src === dst { return true }
}

func main() {
	
	graph := make(map[string][]string)
	graph["f"] = []string{"g", "i"}
	graph["g"] = []string{"h"}
	graph["h"] = []string{}
	graph["i"] = []string{"g", "k"}
	graph["j"] = []string{"i"}
	graph["k"] = []string{}

	hasPath(graph, "f", "j")
}