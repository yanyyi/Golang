package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	ret := integerBreak(2)
	fmt.Println(ret)
	ret = integerBreak(3)
	fmt.Println(ret)
	ret = integerBreak(4)
	fmt.Println(ret)
	ret = integerBreak(5)
	fmt.Println(ret)
	ret = integerBreak(6)
	fmt.Println(ret)
	ret = integerBreak(7)
	fmt.Println(ret)
	ret = integerBreak(8)
	fmt.Println(ret)
	ret = integerBreak(9)
	fmt.Println(ret)
	ret = integerBreak(10)
	fmt.Println(ret)
}
