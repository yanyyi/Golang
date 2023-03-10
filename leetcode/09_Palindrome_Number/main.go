package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rawX := x
	newX := 0
	for x != 0 {
		newX = newX*10 + x%10
		x = x / 10
	}
	if newX == rawX {
		return true
	} else {
		return false
	}
}

//func main() {
//	x := 121
//	fmt.Println(isPalindrome(x))
//}
