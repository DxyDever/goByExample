package main

import "fmt"

func main() {

	var i string

	i = "123"
	/*
	   这么写是有问题的，需要先把i向上转型为接口类型，这样才是比较合适的
	*/
	//switch i.(type) {
	//case int:
	//	fmt.Println("i is a int")
	//case string:
	//	fmt.Println("i is a string")
	//default:
	//	fmt.Println("i I don not know")
	//}

	//这是判断值
	switch i {
	case "123":
		fmt.Println("123")
		break
	default:
		fmt.Println("456")
	}

	/*
		case多个值判断
	*/
	//这是判断值
	switch i {
	case "123", "456":
		fmt.Println("123")
		break
	default:
		fmt.Println("456")
	}

	/*
		空switch
	*/
	switch {
	case i == "123":
		fmt.Println("123")
	default:
		fmt.Println("default")
	}

	//将函数赋给一个变量,只有当这个变量使用()的时候会被调用,记得传入对应的参数
	whatami := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("this is a string")
			break
		case int:
			fmt.Println("this is a int")
			break
		default:
			fmt.Println("default")
		}
	}

	whatami(1)
	whatami("123456789")

}
