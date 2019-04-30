package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var resT []byte
	var flags []int
	doGenereta(n, n, &res, resT, flags)
	return res
}

func doGenereta(n, availabe int, res *[]string, resT []byte, flags []int) {
	if len(resT) == 2*n {
		*res = append(*res, string(resT))
	} else {
		if availabe > 0 {
			doGenereta(n, availabe-1, res, append(resT, '('), append(flags, 0))
			if len(flags) > 0 {
				doGenereta(n, availabe, res, append(resT, ')'), flags[:len(flags)-1])
			}
		} else {
			doGenereta(n, availabe, res, append(resT, ')'), flags[:len(flags)-1])
		}
	}
}
