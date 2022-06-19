package main

import "fmt"

func main() {

	messages := make(chan string)
	//没有设置初始值的大小，所以是不可以添加的
	//messages <- "123"
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	//并没有写进去，而是一直处于阻塞的状态,阻塞就会触发default操作
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
