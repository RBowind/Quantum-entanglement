# ARTS

> ARTS(Algorithm,Review,Tip,Share)

This is my first step in this plan,and this is my first plan in these years. As the saying goes,"Step by Step".So the algorithm's difficulty will become harder and harder.

## 1. Algorithm

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

## 2. Review

I'm learning docker recently, so I would review an article about docker.

As Title says, this article mainly introduces the best practices of Dockerfile and the useful of ENTRYPOINT & CMD, which needed in Dockerfile.As a backend developer,I think Dockerfile is the best way to deploy the environment(development&production).

[Docker ENTRYPOINT & CMD: Dockerfile best practices](https://medium.freecodecamp.org/docker-entrypoint-cmd-dockerfile-best-practices-abc591c30e21)

## 3. Tip

An excellent origin function in PHP:
[`array_multisort()`](http://www.php.net/manual/zh/function.array-multisort.php)

When you want to sort multi-dimensional arrays,this function is exactly what you need!

Just like SQL sort: `order by column1 desc,column2 asc`. You can sort by arrays' specific keys

## 4.Share

[设计模式之意义与原则
](https://github.com/RBowind/RBlog/blob/master/ARTS/week-1/week-1-share.md)
