package main

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDistance(word1 string, word2 string) int {

	n1, n2 := len(word1), len(word2)
	//dp[i][j]存储minDistance(word1[i:],word2[j:])的结果
	dp := make([][]int, n1)
	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {

		}
	}
	if word1[0] == word2[0] {
		return minDistance(word1[1:], word2[1:])
	}
	delete1 := minDistance(word1[1:], word2)
	delete2 := minDistance(word1, word2[1:])
	return min(delete1, delete2)
}

func max(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
