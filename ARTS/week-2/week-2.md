# ARTS-Week2

> ARTS(Algorithm,Review,Tip,Share)

Fine.Second step.

## 1.Algorithm

### Problem(medium)

String to Integer (atoi),Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

- Only the space character ' ' is considered as whitespace character.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

### Solution

Time Complexity: O(n) [code](https://github.com/RBowind/RBlog/blob/master/ARTS/week-2/demo.go)
Use Ascii code and the features of golang string traversal.

```golang
func myAtoi(str string) int {
	var i, num int
	sign := 1
	if str == "" {
		return 0
	}
	// 空格的 ASCII 为 32,把字符串前面的空格过掉
	for i < len(str) && str[i] == 32 {
		i++
	}
	//如果字符串全都是空格
	if i >= len(str) {
		return 0
	}

	//确定数字的正负
	if str[i] == 43 { // + 的ASCII
		i++
	} else if str[i] == 45 { // - 的ASCII
		sign = -1
		i++
	}

	var result int
	for ; i < len(str); i++ {
		// 数字后面遇到了非数字的字符，便停止遍历
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) { // 非 0 ～ 9 的ASCII
			return result
		}

		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		result = num * sign

		if result < math.MinInt32 {
			return math.MinInt32
		} else if result > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return result
}
```

## 2.Review
