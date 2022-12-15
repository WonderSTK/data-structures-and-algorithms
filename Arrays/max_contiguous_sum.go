// Given an array of 􀝊 numbers, give an algorithm for finding
// a contiguous subsequence A(i). . . A(j) for which the sum of elements is maximum.
// Example: {-2, 11, -4, 13, -5, 2} → 20 and {1, -3, 4, -2, -1, 6} → 7
package main

import "fmt"

func MaxContiguousSum(Arr []int) int {
	maxSum, n := 0, len(Arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			currentSum := 0
			for k := i; k <= j; k++ {
				currentSum += Arr[k]
			}
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

func main() {
	Arr :=  []int{-2, 11, -4, 13, -5, 2}
	fmt.Println(MaxContiguousSum(Arr))
}