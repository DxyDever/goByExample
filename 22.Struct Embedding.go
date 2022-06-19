package main

import "fmt"

/*
结构体内嵌
*/

/*

结构体内嵌就类似于继承，类与类之间的继承关系

*/
type base struct {
	num int
}

func (b base) fun() string {
	return fmt.Sprintf("return" + "123")
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "123",
	}

	fmt.Println(co.num, co.str)
	fmt.Println(co.base.num)

	fmt.Println(co.fun())

	type inter interface {
		fun() string
	}

	var bb inter = co

	fmt.Println(bb.fun())
}
