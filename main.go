package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {

	res := []string{}

	var dfs func(prefix string, open int, close int)
	dfs = func(prefix string, open int, close int) {
		if len(prefix) == 2*n {
			res = append(res, prefix)
		}
		if open < n {
			dfs(prefix+"(", open+1, close)
		}
		if close < open {
			dfs(prefix+")", open, close+1)
		}
	}
	dfs("", 0, 0)
	return res
}
