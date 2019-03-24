# ARTS

> ARTS(Algorithm,Review,Tip,Share)

This is my first step in this plan,and this is my first plan in these years. As the saying goes,"Step by Step".So the algorithm's difficulty will become harder and harder.

1. Algorithm

### Problems(easy)

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

#### Example

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

#### Solution

Time Complexity: O(n)

```golang
func twoSum(nums []int, target int) []int {
  m := make(map[int]int)
  for i, n := range nums {
    _, flag := m[n]
    if flag {
      return []int{m[n], i}
      }
      m[target-n] = i
    }
    return nil
}
```

2. Review
