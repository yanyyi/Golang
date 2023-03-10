package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	results := [][]int{}
	//排序
	sort.Ints(nums)
	//三指针法
	for i := 0; i < len(nums)-3; i++ {
		// 跳过重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target2 := target - nums[i]
		//退化为求解三数之和问题
		for j := i + 1; j < len(nums)-2; j++ {
			left, right := j+1, len(nums)-1

			//边界条件
			if nums[j]-target2 > 0 && nums[j] > 0 {
				break
			}

			//跳过重复值
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for left < right {
				if nums[j]+nums[left]+nums[right] == target2 {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					//最小循环去掉重复值
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[j]+nums[left]+nums[right] > target2 {
					right--
				} else {
					left++
				}
			}

		}
	}
	return results

}
