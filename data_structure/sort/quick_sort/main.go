package main

import (
	"fmt"
	"time"
)

func oneQuickSort(arr []int, low int, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high {
			if arr[high] >= pivot {
				high--
				continue
			} else {
				arr[low] = arr[high]
				low++
				break
			}
		}

		for low < high {
			if arr[low] <= pivot {
				low++
				continue
			} else {
				arr[high] = arr[low]
				high--
				break
			}
		}

	}
	arr[low] = pivot
	return low

}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivotpos := oneQuickSort(arr, low, high)
		quickSort(arr, 0, pivotpos-1)
		quickSort(arr, pivotpos+1, high)
	}
}

func main() {
	start := time.Now()
	arr := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("快速排序的结果为:", arr)
	end := time.Since(start)
	fmt.Println("快速排序时间为:", end)
}
