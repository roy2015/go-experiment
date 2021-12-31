package day1

import (
	"fmt"
	"sync"
)

var (
	loopTime  = 100
	waitGroup sync.WaitGroup
)

type printer struct {
	FromChan chan int
	ToChan   chan int
	Key      string
	times    int
}

func (p *printer) print() {
	for true {
		if p.times == loopTime {
			waitGroup.Done()
			return
		}
		<-p.FromChan
		fmt.Print(p.Key)
		p.times++
		p.ToChan <- 1
	}

}

func Main() {
	waitGroup.Add(3)
	chanA := make(chan int, 1)
	chanB := make(chan int, 1)
	chanC := make(chan int, 1)
	Aprinter := printer{FromChan: chanC, ToChan: chanA, Key: "A"}
	Bprinter := printer{FromChan: chanA, ToChan: chanB, Key: "B"}
	Cprinter := printer{FromChan: chanB, ToChan: chanC, Key: "C"}
	chanC <- 10
	go Aprinter.print()
	go Bprinter.print()
	go Cprinter.print()
	//time.Sleep(time.Second * 100)

	waitGroup.Wait()
}
