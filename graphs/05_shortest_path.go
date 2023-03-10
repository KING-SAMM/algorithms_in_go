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


 func shortestPath(edges, nodeA, nodeB) {
	graph := buildGraph(edges)

	queue := make([][]interface{}, 0)
	queue = append(queue, []interface{}{node, 0})
 }

 func main() {
   

// 	const edges = [
//   ['w', 'x'],
//   ['x', 'y'],
//   ['z', 'y'],
//   ['z', 'v'],
//   ['w', 'v']
// ];

// shortestPath(edges, 'w', 'z'); // -> 2


// const edges = [
//   ['w', 'x'],
//   ['x', 'y'],
//   ['z', 'y'],
//   ['z', 'v'],
//   ['w', 'v']
// ];

// shortestPath(edges, 'y', 'x'); // -> 1


// const edges = [
//   ['a', 'c'],
//   ['a', 'b'],
//   ['c', 'b'],
//   ['c', 'd'],
//   ['b', 'd'],
//   ['e', 'd'],
//   ['g', 'f']
// ];

// shortestPath(edges, 'a', 'e'); // -> 3


// const edges = [
//   ['a', 'c'],
//   ['a', 'b'],
//   ['c', 'b'],
//   ['c', 'd'],
//   ['b', 'd'],
//   ['e', 'd'],
//   ['g', 'f']
// ];

// shortestPath(edges, 'e', 'c'); // -> 2



// const edges = [
//   ['a', 'c'],
//   ['a', 'b'],
//   ['c', 'b'],
//   ['c', 'd'],
//   ['b', 'd'],
//   ['e', 'd'],
//   ['g', 'f']
// ];

// shortestPath(edges, 'b', 'g'); // -> -1


// const edges = [
//   ['c', 'n'],
//   ['c', 'e'],
//   ['c', 's'],
//   ['c', 'w'],
//   ['w', 'e'],
// ];

// shortestPath(edges, 'w', 'e'); // -> 1



// const edges = [
//   ['c', 'n'],
//   ['c', 'e'],
//   ['c', 's'],
//   ['c', 'w'],
//   ['w', 'e'],
// ];

// shortestPath(edges, 'n', 'e'); // -> 2



// const edges = [
//   ['m', 'n'],
//   ['n', 'o'],
//   ['o', 'p'],
//   ['p', 'q'],
//   ['t', 'o'],
//   ['r', 'q'],
//   ['r', 's']
// ];

// shortestPath(edges, 'm', 's'); // -> 6

 }