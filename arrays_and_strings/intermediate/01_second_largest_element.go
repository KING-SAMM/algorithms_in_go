package main 

/**
 * Problem: Find the second-largest element in an array.
 * Input: [3, 7, 2, 9, 5]
 * Output: 7
 * Input: [-3, -7, -2, -9, -5]
 * Output: -3
 */

 /**
  * Solution: Finds the second-largest element in an array.
  *
  * @param {Array} arr - The input array.
  * @returns {Integer} - The second-largest element in the array, or -1 if not found.
  */

import "fmt"

func findSecondLargest(arr []int) int {
	largest, secondLargest := larger(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} 

		if arr[i] < largest && arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}

func larger(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func main() {
	arr := []int{-3, -7, -2, -9, -5}
	secLargest := findSecondLargest(arr)
	fmt.Println(secLargest)
}
