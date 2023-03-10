package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	//假设第一个字符最长
	result := string(strs[0])
	//比较每一个字符和最长字符的公共部分
	for i, str := range strs {
		if i > 0 {
			result = commonPrefix(result, string(str))
		}
	}
	//返回结果
	return result
}

// 比较两个字符串的共同前缀
func commonPrefix(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	ret := ""
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] == s2[i] {
			ret += string(s1[i])
			continue
		} else {
			break
		}
	}
	return ret
}
