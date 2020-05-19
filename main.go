package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(arrangeWords("Keep calm and code on"))
}

func arrangeWords(text string) string {
	str := strings.ToLower(text)
	res := strings.Split(str, " ")
	for i := 0; i < len(res) - 1; i++ {
		for j := 0; j < len(res) - 1 - i; j++ {
			if len(res[j]) > len(res[j + 1]) {
				res[j], res[j + 1] = res[j + 1], res[j]
			}
		}
	}
	res[0] = strings.Title(res[0])
	return strings.Join(res, " ")
}
