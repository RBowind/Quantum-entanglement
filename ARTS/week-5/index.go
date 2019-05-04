package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) (results []string) {
	return addParanthesisis("", n, 0)
}

func addParanthesisis(str string, toOpen, toClose int) (results []string) {
	if toOpen == 0 && toClose == 0 {
		return []string{str}
	}
	if toOpen > 0 {
		results = append(results, addParanthesisis(str+"(", toOpen-1, toClose+1)...)
	}
	if toClose > 0 {
		results = append(results, addParanthesisis(str+")", toOpen, toClose-1)...)
	}

	return
}
