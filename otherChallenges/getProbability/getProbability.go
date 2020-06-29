package getProbability

import (
	"math"
	"math/big"
	"strconv"
)

func getProbability(balls []int) float64 {
	sum, a := 0, ""
	for _, v := range balls {
		sum += v
		a += strconv.Itoa(v)
	}

	myMap := make(map[string]string, 0)
	getMap(0, balls, myMap, sum / 2, 0, "")

	res := big.NewFloat(0)
	for k, v := range myMap {
		res.Add(res, compare(k, v))
	}
	res = res.Quo(res, cal(a))
	r, _ := res.Float64()
	return r
}

func getMap(index int, balls []int, myMap map[string]string, sum, count int, res string) {
	if index < len(balls) {
		for j := 0; j <= balls[index]; j++ {
			count += j
			if count < sum {
				temp := res
				res += strconv.Itoa(j)
				getMap(index + 1, balls, myMap, sum, count, res)
				res = temp
				count -= j
			} else if count > sum {
				return
			} else {
				temp := res
				res += strconv.Itoa(j)
				if index >= len(balls) - 1 {
					rev := getReverse(res, balls)
					_, ok1 := myMap[res]
					_, ok2 := myMap[rev]
					if !ok1 && !ok2 {
						myMap[res] = rev
					}
				} else {
					getMap(index + 1, balls, myMap, sum, count, res)
				}
				res = temp
				count -= j
			}
		}
	}
}

func getReverse(str string, balls []int) string {
	num1, _ := strconv.ParseInt(str, 10, 64)
	num := int(num1)
	res := ""
	for i, v := range balls {
		k := int(math.Pow(10, float64(len(balls) - i - 1)))
		n := num / k
		res += strconv.Itoa(v - n)
		num %= k
	}
	return res
}

func compare(s1, s2 string) *big.Float {
	b1, b2, c1, c2 := []byte(s1), []byte(s2), 0, 0
	for i := range b1 {
		if b1[i] =='0' {
			c1++
		}
		if b2[i] == '0' {
			c2++
		}
	}
	if c1 == c2 {
		res := big.NewFloat(1)
		res = res.Mul(cal(s1), cal(s2))
		if s1 == s2 {
			return res
		} else {

			return res.Mul(res, big.NewFloat(2))
		}
	}
	return big.NewFloat(0)
}

func cal(s string) *big.Float {
	res, _ := strconv.ParseInt(s, 10, 64)
	num := int(res)
	sum, a, b := 0, big.NewFloat(1), big.NewFloat(1)
	for num != 0 {
		n := num % 10
		sum += n
		for i := 2; i <= n; i++ {
			a.Mul(a, big.NewFloat(float64(i)))
		}
		num /= 10
	}
	for i := 2; i <= sum; i++ {
		b.Mul(b, big.NewFloat(float64(i)))
	}
	return big.NewFloat(0).Quo(b, a)
}