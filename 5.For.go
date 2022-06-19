package main

import "fmt"

func main() {

	i := 1

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("123")
		break
	}

	/*
		for循环的各种形式
	*/
}
