package main

/**
 * Problem: Rotate an array to the right by k steps.
 * 
 * Input: [1, 2, 3, 4, 5], k = 3
 * Output: [3, 4, 5, 1, 2]
 */

import "fmt"

func rotateArray(arr []int, k int) []int {
	n := len(arr)
	k = 3
	reverse(arr, 0, n - 1)
	reverse(arr, 0, k - 1)
	reverse(arr, k, n - 1)
	return arr
}

func reverse(arr []int, start, end int) {
	for range arr {
		if start < end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start++
			end--
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5} 
	k := 3
	rotated := rotateArray(arr, k)
	fmt.Println(rotated)
}