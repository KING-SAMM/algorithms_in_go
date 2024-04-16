package main

/**
 * Problem: Remove duplicate elements from an array.
 * 
 * Input: [1, 2, 2, 3, 4, 4, 5]
 * Output: [1, 2, 3, 4, 5]
 */


import (
    "fmt"
)

func removeDuplicates(slice []int) []int {
    encountered := map[int]bool{}
    result := []int{}

    for _, v := range slice {
        if encountered[v] == false {
            encountered[v] = true
            result = append(result, v)
        }
    }

    return result
}

func main() {
    // Define a slice with duplicate elements
    numbers := []int{1, 2, 2, 3, 4, 4, 5}

    // Remove duplicates from the slice
    uniqueNumbers := removeDuplicates(numbers)

    // Print the result
    fmt.Println("Original Slice:", numbers)
    fmt.Println("Slice with Duplicates Removed:", uniqueNumbers)
}