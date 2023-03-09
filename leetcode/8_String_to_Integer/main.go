package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	//去掉空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		sign = -1
	}
	//处理数字
	var result = 0
	for _, ch := range s {
		if ch > '9' || ch < '0' {
			break
		}
		result = result*10 + int(ch-'0')
		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if result*sign < math.MinInt32 {
			return math.MinInt32
		}
	}

	return result * sign
}

//func main() {
//	s := " -12654 kok"
//	fmt.Println(myAtoi(s))
//}
