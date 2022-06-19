package main

import "fmt"

func main() {

	i := 0

	if i == 10 {
		fmt.Println(i)
	} else {
		fmt.Println("10")
	}

	/*
		对于if 语句而言，我们还可以在判断之前加上赋值语句
	*/
	if num := 9; num <= 0 {
		fmt.Println("num 小于 0")
	} else if num > 10 {
		fmt.Println("num 大于 10")
	} else {
		fmt.Println(num)
	}

}
