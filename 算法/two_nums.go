package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, prs := m[n]
		fmt.Println(m[n])
		fmt.Println(prs)
		if prs {
			return []int{m[n], i}
		}
		m[target-n] = i
		fmt.Println(m)
	}
	return nil
}

func main() {
	nums := []int{1, 7, 3, 4, 5}
	result := twoSum(nums, 6)
	fmt.Println(result)
}
