package main

import "fmt"

func main() {

	/*
		两种创建map的方式
	*/
	map1 := make(map[string]int)

	map1["1"] = 1

	map1["2"] = 2

	delete(map1, "1")

	fmt.Println(map1)

	map2 := map[string]int{
		"3": 3,
		"4": 4,
	}

	fmt.Println(map2)

	for k, v := range map2 {
		fmt.Println(k, v)
	}

}
