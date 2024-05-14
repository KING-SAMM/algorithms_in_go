package main

/**
 * Problem: Merge two sorted arrays
 * inputs: arr1 = [1, 3, 5, 7], arr2 = [2, 4, 6, 8];
 * Output: [1, 2, 3, 4, 5, 6, 7, 8]
 */
import "fmt"

 func mergeSortedArrays(arr1, arr2 []int) []int {
	resultArr := []int{}
	i, j := 0, 0 // Indices for arr1 and arr2

	// Iterate through both arrays
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultArr = append(resultArr, arr1[i])
			i++
		} else {
			resultArr = append(resultArr, arr2[j])
			j++
		}
	}
	
	 // Add remaining elements from arr1, if any
	for i < len(arr1) {
		resultArr = append(resultArr, arr1[i])
		i++
	}

	 // Add remaining elements from arr2, if any
	for j < len(arr2) {
		resultArr = append(resultArr, arr2[j])
		j++
	}

	return resultArr
 }


 func main() {
	 arr1 := []int{1, 3, 5, 7}
	 arr2 := []int{2, 4, 6, 8}

	 fmt.Println(mergeSortedArrays(arr1, arr2))
 }