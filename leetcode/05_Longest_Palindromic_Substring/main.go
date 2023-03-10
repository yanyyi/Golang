package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := string(s[0])
	maxLen := 0
	for low := 0; low < len(s); low++ {
		if len(s)-low < maxLen {
			break
		}
		for high := low + 1; high < len(s)+1; high++ {
			if isPalindrome(s[low:high]) {
				if high-low > maxLen {
					result = s[low:high]
					maxLen = high - low
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	return result

}

func isPalindrome(s string) bool {
	llen := len(s)
	for i := 0; i < llen/2; i++ {
		if s[i] != s[llen-1-i] {
			return false
		}
	}
	return true
}

//func main() {
//	s := "abdoa"
//	fmt.Println(isPalindrome(s))
//}
