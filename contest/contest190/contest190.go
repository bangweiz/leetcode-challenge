package contest190

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	a := strings.Split(sentence, " ")
	for i := 0; i < len(a); i++ {
		if strings.HasPrefix(a[i], searchWord) {
			return i + 1
		}

	}
	return -1
}

func maxVowels(s string, k int) int {
	a := []byte(s)
	max := 0
	res := 0
	pointer1 := 0
	pointer2 := k
	for i := pointer1; i < pointer2; i++ {
		if isVowels(a[i]) {
			max++
		}
	}
	res = max
	for pointer2 < len(a) {
		res1, res2 := isVowels(a[pointer1]), isVowels(a[pointer2])
		if res1 {
			if !res2 {
				max--
			}
		} else {
			if res2 {
				max++
				if res < max {
					res = max
				}
			}
		}
		pointer2++
		pointer1++
	}
	return res
}

func isVowels(a byte) bool {
	return a  == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pseudoPalindromicPaths (root *TreeNode) int {
	count := []int{0}
	myMap := make(map[int]int)
	dfs(root, myMap, count)
	return count[0]
}

func dfs(root *TreeNode, myMap map[int]int, count []int) {
	if v, ok := myMap[root.Val]; ok {
		myMap[root.Val] = v + 1
	} else {
		myMap[root.Val] = 1
	}
	if root.Left != nil {
		dfs(root.Left, myMap, count)
	}
	if root.Right != nil {
		dfs(root.Right, myMap, count)
	}
	if root.Right == nil && root.Left == nil {
		flag := 0
		for _, v := range myMap {
			if v % 2 != 0 {
				flag++
			}
		}
		if flag < 2 {
			count[0]++
		}
	}
	myMap[root.Val] = myMap[root.Val] - 1
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	data := make([][]int, len(nums1))
	for i := range data {
		data[i] = make([]int, len(nums2))
	}
	for i := range nums1 {
		for j := range nums2 {
			product := nums1[i] * nums2[j]

			if i == 0 {
				if j == 0 {
					data[i][j] = product
				} else {
					if product > data[i][j - 1] {
						data[i][j] = product
					} else {
						data[i][j] = data[i][j - 1]
					}
				}
			} else if j == 0 && i > 0 {
				if product > data[i - 1][j] {
					data[i][j] = product
				} else {
					data[i][j] = data[i - 1][j]
				}
			} else {
				if data[i - 1][j - 1] > 0 {
					product += data[i - 1][j - 1]
				}
				if product >= data[i][j - 1] && product >= data[i - 1][j] {
					data[i][j] = product
				} else {
					if data[i - 1][j] > data[i][j - 1] {
						data[i][j] = data[i - 1][j]

					} else {
						data[i][j] = data[i][j - 1]
					}
				}
			}
		}
	}
	return data[len(nums1) - 1][len(nums2) - 1]
}