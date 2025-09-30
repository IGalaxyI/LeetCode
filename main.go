package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := []string{}
	var dfs func(pos int, prefix string)

	dfs = func(pos int, prefix string) {
		if pos == len(digits) {
			res = append(res, prefix)
			return
		}
		letters := mapping[digits[pos]]
		for i := 0; i < len(letters); i++ {
			dfs(pos+1, prefix+letters[i:i+1])
		}
	}
	dfs(0, "")
	return res

}
