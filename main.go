package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman2(3749))
}

func intToRoman(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	th := num / 1000
	h := (num % 1000) / 100
	t := (num % 100) / 10
	o := num % 10

	result := thousands[th] + hundreds[h] + tens[t] + ones[o]
	return result
}

func intToRoman2(num int) string {
	type pair struct {
		val int
		sym string
	}

	pairs := []pair{{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"}}

	var res strings.Builder
	for _, p := range pairs {
		for num >= p.val {
			res.WriteString(p.sym)
			num -= p.val
		}
	}
	return res.String()
}
