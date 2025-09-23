package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(9646324351))
}

func reverse(x int) int {
	const MaxInt32 = 1<<31 - 1
	const MinInt32 = -1 << 31
	save := 0
	if x < 0 {
		save = x
		x = -x
	}
	s := 0
	for x > 0 {
		d := x % 10
		if s > (MaxInt32-d)/10 {
			return 0
		}
		s = s*10 + d
		x /= 10
	}
	if save < 0 {
		return -s
	} else {
		return s
	}
}
