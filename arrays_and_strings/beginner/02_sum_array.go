package main

/**
 * Problem: 
 * Calculate the sum of all numbers in an array.
 * 
 * Input: [2, 4, 6, 8, 10]
 * Output: 30
 */

 import "fmt"

 func calculateSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
 }

 func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := calculateSum(arr)
	fmt.Println(sum)
 }