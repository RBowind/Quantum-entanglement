package main

import (
	"fmt"
	"math"
	"strconv"
)

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

	for ; i < len(str); i++ {
		// 数字后面遇到了非数字的字符，便停止循环
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) { // 非 0 ～ 9 的ASCII
			return num * sign
		}
		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		if num*sign < math.MinInt32 {
			return math.MinInt32
		} else if num*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return num * sign
}

func main() {

	stringArray := []string{"42", "  42", "  -13", "array-13", "813 31is my birthday813", "91283472332"}

	for _, val := range stringArray {
		fmt.Println(myAtoi(val))
	}
}
