package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	sum := 0

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum = nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}

			}
		}

	}
	return result
}
