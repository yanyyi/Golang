package main

import "sort"

func threeSum(nums []int) [][]int {
	//初始化一个二维数组
	result := [][]int{}

	sort.Ints(nums)
	// 双指针法
	llen := len(nums)
	for i := 0; i < llen-2; i++ {
		//初始化索引值
		left, right := i+1, llen-1

		//边界条件
		if nums[i] > 0 {
			break
		}

		//跳过重复值(大循环)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			// 双指针找数
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//跳过重复值(小循环)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}

	}

	return result
}
