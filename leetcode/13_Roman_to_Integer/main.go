package main

func romanToInt(s string) int {
	llen := len(s)
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := romanMap[s[llen-1]]
	//从倒数第二位 llen-2开始往前推到 0, 当前位比后一位大或者相等，加上当前位； 当前位比后一位小，减掉当前位
	for i := llen - 2; i >= 0; i-- {
		if romanMap[s[i]] >= romanMap[s[i+1]] {
			result += romanMap[s[i]]
		} else {
			result -= romanMap[s[i]]
		}
	}
	return result
}
