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

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 == 1 {
		return float64(nums1[n/2])
	} else {
		m1 := nums1[n/2-1]
		m2 := nums1[n/2]
		return float64(m1+m2) / 2.0
	}
}
