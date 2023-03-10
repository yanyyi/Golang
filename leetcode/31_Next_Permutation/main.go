package main

import (
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	var i int

	//原先完全降序
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			continue
		} else {
			break
		}
	}
	if i == len(nums)-1 {
		for j := 0; j < len(nums)/2; j++ {
			nums[j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[j]
		}
		return
	}

	//原先没有完全降序
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		} else {
			//找到比nums[i]不小于的数中最小的那个
			min_i := greaterMin(nums, i)
			//交换这两个数
			nums[i], nums[min_i] = nums[min_i], nums[i]
			//nums[i+1:]降序排列
			sort.Ints(nums[i+1:])
			break
		}
	}
}

func greaterMin(nums []int, index int) int {
	min := nums[index+1]
	min_i := index + 1
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] && nums[i] < min {
			min = nums[i]
			min_i = i
		}
	}
	return min_i

}

//func main() {
//	arr := []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
//	nextPermutation(arr)
//	fmt.Println(arr)
//}
