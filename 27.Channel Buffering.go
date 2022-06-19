package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan string, 2)

	message <- "1"
	message <- "2"

	go func(chan string) {
		time.Sleep(time.Second * 3)
		for {
			msg := <-message
			fmt.Println(msg)
		}
	}(message)

	time.Sleep(time.Second * 5)
}
