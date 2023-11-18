package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	var i int
	for i = 2; i < n+1; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[i-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
