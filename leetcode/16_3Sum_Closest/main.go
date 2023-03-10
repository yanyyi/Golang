package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for {
			if start < end {
				sum := nums[i] + nums[start] + nums[end]
				if abs(sum-target) < abs(result-target) {
					result = sum
				}
				if sum < target {
					start++
				} else if sum > target {
					end--
				} else {
					return sum
				}
			} else {
				break
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
