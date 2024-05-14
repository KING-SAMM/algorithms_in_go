package main

/**
 * Problem: Find the index of a specific element in an array
 * 
 * Input: [3, 5, 2, 7, 1], element = 7
 * Output: 3
 */

/**
 * Solution: Finds the index of an element in an array.
 *
 * @param {[]int} arr - The input array.
 * @param {Integer} element - The element to search for.
 * @returns {Integer} - The index of the element in the array, or -1 if not found.
 */

import "fmt"

func findIndex(arr []int, element int) int {
	for i, _ := range arr {
		if arr[i] == element {
			return i
		}
	}

	return -1
}

func main() {
	arr := []int{3, 5, 2, 7, 1}
	el := 4
	index := findIndex(arr, el)
	fmt.Println(index)
}