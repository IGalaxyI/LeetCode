package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	s := float64(0)
	n := float64(len(nums1))
	if len(nums1)%2 == 1 {
		return float64(nums1[n/2])
	} else {
		m1 := nums1[n/2-1]
		m2 := nums1
		return float64(r)
	}
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	print("момент")
// }
