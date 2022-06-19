package main

import "fmt"

/*
闭包
*/

func seq() func() int {

	i := 0

	return func() int {
		i++
		return i
	}
}
func main() {

	fun2 := seq()
	fmt.Println(fun2())
	fmt.Println(fun2())

	fun2 = seq()
	fmt.Println(fun2())

}
