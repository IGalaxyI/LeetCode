package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for k := n - 1; k >= 2; k-- {
		left, right := 0, k-1
		for left < right {
			if nums[left]+nums[right] > nums[k] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}
