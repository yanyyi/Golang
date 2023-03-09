package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	llen := len(nums1) + len(nums2)
	s := make([]int, llen)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			s[i+j] = nums1[i]
			i++
			continue
		} else {
			s[i+j] = nums2[j]
			j++
			continue
		}
	}

	//若nums1遍历完全
	if i == len(nums1) {
		copy(s[i+j:], nums2[j:])

		//判断数组长度奇偶数情况
		if len(s)%2 == 1 {
			return float64(s[len(s)/2])
		} else {
			sum := s[len(s)/2] + s[len(s)/2-1]
			return float64(sum) / 2
		}
	}

	//若nums2遍历完全
	if j == len(nums2) {
		copy(s[i+j:], nums1[i:])

		//判断数组长度奇偶数情况
		if len(s)%2 == 1 {
			return float64(s[len(s)/2])
		} else {
			sum := s[len(s)/2] + s[len(s)/2-1]
			return float64(sum) / 2
		}
	}
	return 0
}

//func main() {
//	nums1 := []int{1, 2}
//	nums2 := []int{3, 4}
//	ret := findMedianSortedArrays(nums1, nums2)
//	fmt.Println(ret)
//}
