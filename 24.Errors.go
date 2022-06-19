package main

import (
	"errors"
	"fmt"
)

//接口类型可以接受任意变量
func f1(arg int) (int, error) {

	if arg == 42 {
		return -1, errors.New("can not work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

//由于它实现了这个Error方法，重写了，就表示这个类实现了接口Error
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//用父类接口变量来接收这个实现了父类接口的实体对象，&argError{arg, "can not work with it"}
func f2(arg int) (int, error) {
	if arg == 42 {
		//接口父变量接收的是引用类型的
		return -1, &argError{arg, "can not work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//这里返回的e的类型是error接口类型，所以还需要进行.(type)进行转型
	_, e := f2(42)
	//fmt.Println(e.arg)
	//fmt.Println(e.prob)
	//强制把e转型为子类的类型ae,是一个指针类型
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	/*
		上面的实现都是基于引用传递的，其实还可以写成值传递的,引用传递会更加的节约内存的
	*/

	//由于它实现了这个Error方法，重写了，就表示这个类实现了接口Error
	//func (e argError) Error() string {
	//	return fmt.Sprintf("%d - %s", e.arg, e.prob)
	//}
	//
	////用父类接口变量来接收这个实现了父类接口的实体对象，&argError{arg, "can not work with it"}
	//func f2(arg int) (int, error) {
	//	if arg == 42 {
	//		//接口父变量接收的是引用类型的
	//		return -1, argError{arg, "can not work with it"}
	//	}
	//	return arg + 3, nil
	//}
	//
	//
	//if ae, ok := e.(argError); ok {
	//	fmt.Println(ae.arg)
	//	fmt.Println(ae.prob)
	//}

}
