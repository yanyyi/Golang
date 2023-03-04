package main

import "fmt"

func BinarySearchLast(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			if mid == len(arr)-1 || arr[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{7, 33, 33, 33, 33, 37, 41, 43}
	target := 33
	i := BinarySearchLast(arr, target)
	fmt.Println("target num last occur: i=", i)

}
