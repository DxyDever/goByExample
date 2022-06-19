package main

import "fmt"

/*
只读，只写channel
*/

//这里的pings是只写，限制方法中渠道的方向性
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//这里的pings是只读，pong是致谢
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}