package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
func main() {

	done := make(chan bool, 1)

	go worker(done)

	//主协程等待从协程完成任务之后再结束
	<-done
}
