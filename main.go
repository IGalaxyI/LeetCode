package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotal([][]int{[[2],[3,4],[6,5,7],[4,1,8,3]]}))
}

func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0{
		return triangle[0][0]
	}
	
	dp := make([]int, len(triangle)+1)

	for i := len(triangle)-1; i>=0;i--{
		for j := 0; j<len(triangle[i]);j++{
			dp[j] = triangle[i][j] + min(dp[j],dp[j+1])
		}
	}
	return dp[0]

}