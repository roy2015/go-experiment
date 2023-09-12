package main

//测试多goroutine并发下的range chan是否重复轮询，发现是隔离的，goroutine a range了的chan里的某数据，
//goroutine b就不会再重复range
//读写都是异步的
import (
	"fmt"
	"strconv"
	"time"
)

func writeChan(write_chan chan<- string, idx int, msg string) {
	fmt.Printf("write routine:  idx: %d msg: %s \n", idx, msg)
	write_chan <- msg
}

func testChan(idx int, str_stream <-chan string) {
	for str := range str_stream {
		fmt.Printf("----- reader routine: idx: %d msg %s \n", idx, str)
	}
}

func main() {
	str_chan := make(chan string)

	defer close(str_chan)

	//write 20 times
	for i := 0; i < 20; i++ {
		go writeChan(str_chan, i, strconv.Itoa(i))
	}

	//10 goroutine loop
	for i := 0; i < 10; i++ {
		go testChan(i, str_chan)
	}

	//阻塞等待，十秒后解除
	done := make(chan bool)
	time.AfterFunc(time.Second*10, func() {
		done <- true
	})
	<-done

}
