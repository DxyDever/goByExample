package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	/*
		一共有50个协程去帮忙做这件事情
	*/
	for i := 0; i < 50; i++ {

		wg.Add(1)

		go func() {
			for c := 0; c < 100; c++ {
				atomic.AddUint64(&ops, 2)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops", ops)

}
