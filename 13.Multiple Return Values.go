package main

import "fmt"

func fun() (int, int) {

	return 2, 2
}

func main() {

	a, b := fun()

	fmt.Println(a, b)
	
}
