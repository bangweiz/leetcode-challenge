package main

import (
	"fmt"
	"github/bangweiz/leetCode/june_challenge"
)

func main() {
	people := make([][]int, 6)
	people[0] = []int{7,0}
	people[1] = []int{4,4}
	people[2] = []int{7,1}
	people[3] = []int{5,0}
	people[4] = []int{6,1}
	people[5] = []int{5,2}

	res := june_challenge.ReconstructQueue(people)

	for _, v := range res {
		fmt.Printf("%v %v \n", v[0], v[1])
	}
}
