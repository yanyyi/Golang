package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	low, high := 0, 0
	maxLen := 0
	m := make(map[byte]int)
	for high < len(s) {
		char := s[high]
		m[char]++
		//遇到重复的字符(窗口内有2个),左指针移动到第1个重复字符的右边
		for m[char] > 1 {
			m[s[low]]--
			low++
		}

		//比较现有窗口和最大窗口
		maxLen = max(maxLen, high-low+1)
		//右指针右移
		high++
	}
	return maxLen
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
