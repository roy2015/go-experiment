package main

import (
	"fmt"
	"time"
)

//测试进度条......
func main() {
	quit := make(chan bool)
	fmt.Printf("Download Test: ")
	go dots(quit)
	time.Sleep(time.Second * 5)
	quit <- true
}

//两百毫秒走一格
func dots(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Millisecond * 200)
			fmt.Print(".")
		}
	}
}
