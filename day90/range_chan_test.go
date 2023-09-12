package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRangeChan(t *testing.T) {
	dataChan := make(chan int)

	//单独线程 range channel
	go func() {
		for i := range dataChan {
			fmt.Printf("%d\n", i)
		}
	}()

	//延迟写
	time.AfterFunc(time.Second*1, func() {
		for i := 0; i < 100; i++ {
			dataChan <- i
		}
	})

	//阻塞等待，十秒后解除
	done := make(chan bool)
	time.AfterFunc(time.Second*10, func() {
		close(dataChan)
		done <- true
	})
	<-done

}
