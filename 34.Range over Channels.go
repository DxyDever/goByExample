package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//只可以直接遍历值，不可以得到对应的下标，没有下标的概念
	for element := range queue {
		fmt.Println(element)
	}

}
