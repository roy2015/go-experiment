package main

import (
	"fmt"
	"math"
)

/*
执行结果：通过

执行用时：12 ms, 在所有 Go 提交中击败了47.73%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了56.82%的用户
通过测试用例：93 / 93
*/
func minFlipsMonoIncr(s string) int {
	bytes := []byte(s)
	if len(bytes) == 1 {
		return 0
	}

	//第一位：翻转后的值，第二位：是否翻转  eg:n11 0翻转成1  n00:未翻转0
	var preN0 int
	var preN1 int
	//先求第一个n1, n0（preN1, preN0）
	first := bytes[0]
	if first == '0' {
		preN1 = 1
	} else {
		preN0 = 1
	}
	var n1 int
	var n0 int
	for idx, b := range bytes {
		if idx == 0 {
			continue
		}
		//当前是'0'
		if b == '0' {
			n1 = int(math.Min(float64(preN1+1), float64(preN0+1)))
			n0 = preN0
		} else {
			n1 = int(math.Min(float64(preN0), float64(preN1)))
			n0 = preN0 + 1
		}
		preN0 = n0
		preN1 = n1
	}
	return int(math.Min(float64(n1), float64(n0)))
}

func main() {
	fmt.Println(minFlipsMonoIncr("0110"))     //1
	fmt.Println(minFlipsMonoIncr("010110"))   //2
	fmt.Println(minFlipsMonoIncr("00011000")) //2
}
