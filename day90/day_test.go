package main

//测试函数返回值，只要return不用

import (
	"fmt"
	"testing"
)

func fun1() (n uint32, s string) {
	n = 20
	s = "roy"
	return
}

func TestFuncRet(t *testing.T) {
	n, s := fun1()
	fmt.Println("n: ", n, "  s: ", s)
}

func TestMap(t *testing.T) {
	my_map := map[int](string){}
	my_map[1] = "guo"
	my_map[11] = "jun"
	key := 2
	if val, isOk := my_map[key]; isOk {
		fmt.Println("key ", key, " value:  ", val)
	} else {
		fmt.Println("key is not exist ")
	}
}
