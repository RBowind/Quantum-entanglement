# ARTS-Week3


## Algorithm
**3 Sum**
### problem:
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

> Note:The solution set must not contain duplicate triplets.

### Solution:
```golang
func threeSum(nums []int) [][]int {
    length := len(nums)
	solution := make([][]int, 0)	
	sort.Ints(nums)
	for k, _ := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l := k + 1
		r := length-1
		for l < r {
			sum := nums[k] + nums[l] + nums[r]
			if sum == 0 {
				array := []int{nums[k], nums[l], nums[r]}
				solution = append(solution, array)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}

		}
	}
	return solution
}
```



## Review
[Golang — handling errors gracefully](https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f) 

Handling errors in golang is special and unique.In this article, You can get lots of details about the golang errors handling skill. What's more, you should clear what error infomation is exactly what you need.


## Tip
I want to share my opinion that when should we use "panic" in golang.
We all know that "panic" would lead to the program stop immediately.
Therefore, most of the time, we won't use it. But I still list two situation that we can use it:
- The program is in an unrecoverable state which means it would lead to more problems if it still execute.
- An unhandled error has occurred.


## Share

I left office last week. As my charge said,"If you can't grow up and get paid enough,it's time to leave." So I told my resign plan to my charge ahead of time,and I think the handover needs one week at least. However,after the company failed to persuade me to stay,they just gave me one day to prepare the handover.Maybe it's time to pay for my social security and something else,or thought that I would be in terrible work situation during handover time,or whatever.

No matter what it is,the last day was a busy day which made me full. Also,I felt nervous cause I thought this was a unprepared handover,means a bad ending.

Yes,I feel nervous up to now.Hope they have a happy starting and pay more attention to the tech team.

And these weeks, I felt too terrible to keep my ARTS plan. Now, I come back and get my hands working.

Now I just want to reflect myself. Why can't I hand over my work in one day.
1. I didn't picture my logic in every micro system that I maintained.
2. I didn't write my work methods to group document.
3. I shouldn't use other program language to achieve the demands.
4. It was too casual when submiting code. 
5. I didn't optimize my code struct when I was idle.

