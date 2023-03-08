package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, ok := m[target-num]
		if ok {
			return []int{j, i}
		} else {
			m[num] = i
		}
	}
	return []int{}
}
