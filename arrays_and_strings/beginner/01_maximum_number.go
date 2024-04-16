package main

/**
 * Problem: Find the maximum number in an array.
 * Input: [4, 9, 2, 7, 5]
 * Output: 9
 */

/**
 * Solution: 
 * maxNum finds the maximum value in an array.
 *
 * @param {[]int} arr - The input array.
 * @returns {int} max - The maximum value found in the array.
 */
import "fmt"

func maxNum(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr) - 1; i++{
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := []int{3, 8, 12, 9, 8, -4}
	max := maxNum(arr)
	fmt.Println(max)
}