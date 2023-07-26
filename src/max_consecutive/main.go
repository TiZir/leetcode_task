package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left, curr, buf, max := 0, 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}
		for curr > k {
			if nums[left] == 0 {
				curr--
			}
			buf--
			left++
		}
		buf++
		if buf > max {
			max = buf
		}
	}
	return max
}

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Printf("%d", longestOnes(arr, 2))
}
