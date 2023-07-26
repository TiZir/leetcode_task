package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	var maxAvg float64
	for i := 0; i < k; i++ {
		maxAvg += float64(nums[i])
	}
	curr := maxAvg
	for i := k; i < len(nums); i++ {
		curr += float64(nums[i]) - float64(nums[i-k])
		if curr > maxAvg {
			maxAvg = curr
		}
	}
	maxAvg /= float64(k)
	return maxAvg
}

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	fmt.Printf("%.5f", findMaxAverage(arr, 4))
}
