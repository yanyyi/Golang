package main

func removeDuplicates(nums []int) int {
	llen := len(nums)
	count := 0
	for i := 1; i < llen; i++ {
		if nums[i] == nums[i-1] {
			count++
			continue
		} else {
			nums[i-count] = nums[i]
		}
	}
	return llen - count
}
