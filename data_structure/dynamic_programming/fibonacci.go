package main

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	num1 := 1
	num2 := 1
	var result = 0
	for i := 0; i < n-2; i++ {
		result = num1 + num2
		num1, num2 = num2, result
	}
	return result
}
