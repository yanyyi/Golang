package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	results := []int{-1, -1}
	low := 0
	high := len(nums) - 1
	mid := -1
	//寻找第一次出现的位置
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				results[0] = mid
				break
			} else {
				high = mid - 1
			}
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	//寻找最后一次出现的位置
	if results[0] == -1 {
		return []int{-1, -1}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			results[1] = j
			break
		}

	}
	return results

}
