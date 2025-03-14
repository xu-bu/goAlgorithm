package main

import "fmt"

//416
//将问题抽象为价值等于重量，容量为sum//2的背包的最大价值是否为sum//2
func canPartition(nums []int) bool {
	bagSize := 0
	for _, v := range nums {
		bagSize += v
	}
	if bagSize%2 == 1 {
		return false
	}
	bagSize /= 2

	dp := make([]int, bagSize+1)
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}

		}
	}
	return dp[bagSize] == bagSize
}

func findTargetSumWays(nums []int, target int) int {
	//dp[i][j]表示nums[:i+1],target=j的结果
	//dp[i][j]=dp[i-1][j+nums[i]]+dp[i-1][j-nums[i]]
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum || target < 0-sum {
		return 0
	}
	size := sum * 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, size+1)
	}
	dp[0][sum-nums[0]] += 1
	dp[0][sum+nums[0]] += 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		for j := 0; j <= size; j++ {
			if j+v <= size {
				dp[i][j] += dp[i-1][j+v]
			}
			if j-v >= 0 {
				dp[i][j] += dp[i-1][j-v]
			}

		}
	}
	return dp[len(nums)-1][target+sum]
}
