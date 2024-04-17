package  main
/**
 * Problem: 
 * Check if an array is sorted in ascending order.
 * 
 * Input: [2, 4, 6, 8, 10]
 * Output: true
 */

 /**
 * Solution: Checks if the array is sorted in ascending order
 * 
 * @param {[]int} arr - The input array
 * @returns {Bool} - The indicator whether the array is sorted
 */

 import "fmt"

 func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			return false
		}
	}
	return true
 }

 func main() {
	arr := []int{2, 4, 6, 8, 10}

	issorted := isSorted(arr)
	fmt.Println(issorted)
 }

