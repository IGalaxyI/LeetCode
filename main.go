package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 5))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	left := 0
	right := 0
	sum := 0
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
