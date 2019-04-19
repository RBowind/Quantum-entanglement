package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	res := make([][]int, 0)

	if len(nums) < 4 {
		return nil
	}

	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			return [][]int{nums}
		}
		return nil
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		for j := i + 1; j < len(nums)-2; j++ {
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// fmt.Println(i, j, l, r)
					// fmt.Println(res)
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}

			}
			//排除重复的情况
			for j < len(nums)-2 && nums[j+1] == nums[j] {
				j++
			}
		}
		// 排除重复的情况
		for i < len(nums)-1 && nums[i+1] == nums[i] {
			fmt.Println(i)
			i++
		}

	}
	return res
}

func main() {

	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res)
}
