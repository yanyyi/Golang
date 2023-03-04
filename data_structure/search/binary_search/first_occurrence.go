package main

import "fmt"

func BinarySearchFirst(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0
	for low <= high {
		//mid = (low + high) / 2
		mid = low + (high-low)/2
		if target == arr[mid] {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{7, 33, 33, 33, 33, 37, 41, 43}
	target := 33
	i := BinarySearchFirst(arr, target)
	fmt.Println("target num first occur: i=", i)

}
