package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	var mid int = 0
	for low <= high {
		//mid = (low + high) / 2
		mid = low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1

}

func main() {
	arr := []int{7, 10, 13, 16, 19, 29, 32, 33, 37, 41, 43}
	target := 33
	i := BinarySearch(arr, target)
	fmt.Println("i=", i)

}
