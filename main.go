package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("-042"))
}

func myAtoi(s string) int {
	const MinInt32 = -1 << 31
	const MaxInt32 = 1<<31 - 1

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	sign := 1

	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	res := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if res > (MaxInt32-digit)/10 {
			if sign == 1 {
				return MaxInt32
			}
			return MinInt32
		}
		res = res*10 + digit
		i++
	}

	return sign * res
}
