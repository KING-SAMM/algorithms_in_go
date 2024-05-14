package main

/**
 * Problem: Reverse the elements of an array
 * input: [1, 2, 3, 4, 5]
 * Output: [5, 4, 3, 2, 1]
 */

 /**
 * Solution: Reverse the elements of an array
 * 
 * @param {Array} arr - The input array
 * @param {integer} start - The index of the first element
 * @param {integer} end - The index of the last element
 */

 import "fmt"

//  func reverse(arr []int) []int {
// 	length := len(arr)/2
// 	start := 0
// 	end := len(arr) - 1
// 	for i := 0; i <= length; i++ {
// 		if start < end {
// 			temp := arr[start]
// 			arr[start] = arr[end]
// 			arr[end] = temp
// 			start++
// 			end--
// 		}
// 	}
// 	return arr
//  }

 func reverse(a []int) []int {
	b := []int{}
	for i := len(a) - 1; i >= 0; i-- {
        b = append(b, a[i])
    }
	return b
 }

 func main() {
	arr := []int{1, 2, 3, 4, 5}
	reversed := reverse(arr)
	fmt.Println(reversed)
 }