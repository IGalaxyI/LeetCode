package main

import (
	"sort"
	"strings"
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

func convert(s string, numRows int) string {
	s := "PAYPALISHIRING"
	numRows := 3

	if numRows <= 1 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	row := 0
	for _, i := range s {
		rows[row].WriteRune(i)
		if row == 0 {
			dir = 1
		} else if row == numRows-1 {
			dir = -1
		}
		row += dir
	}
	var out strings.Builder
	for i := 0; i < numRows; i++ {
		out.WriteString(rows[i].String())
	}
	return out.String()
}
