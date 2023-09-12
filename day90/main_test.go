package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBinary(t *testing.T) {
	fmt.Println("整型转成2进制", strconv.FormatInt(100, 2))
	i, _ := strconv.ParseInt("100", 2, 0)
	fmt.Println("字符串转数字", i)
	fmt.Println("或运算", 1|2)
}

func Test_array_append(t *testing.T) {
	var array = make([]int, 1)
	array[0] = 3
	fmt.Println("数组append ", array)
}

// make函数数组，append函数定义
func Test_fuc_append(t *testing.T) {
	var arrayFunc = make([]func(a1 int, a2 int) int, 0)
	arrayFunc = append(arrayFunc, func(a1 int, a2 int) int {
		return a1 + a2
	}, func(a1 int, a2 int) int {
		return a1 * a2
	})
	for _, f := range arrayFunc {
		fmt.Println("数组append ", f(9, 10))
	}
}
