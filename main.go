package main

import "fmt"

func main() {
	power, res := []int{1}, 0
	for i := 1; i <= 64; i++ {
		power = append(power, power[i - 1] * 2 % 1000000007)
	}
	for i := 0; i < 64; i++ {
		res = (res + power[i]) % 1000000007
	}
	fmt.Println(res)
}
