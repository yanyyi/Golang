package main

func maxArea(height []int) int {
	// 双指针法
	result := 0
	low := 0
	high := len(height) - 1
	for low < high {
		if (high-low)*min(height[low], height[high]) > result {
			result = (high - low) * min(height[low], height[high])
		}
		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return result

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
