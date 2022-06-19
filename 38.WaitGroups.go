package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
}

/*
Go语言标准库中WaitGroup只有三个方法：

Add(delta int)表示向内部计数器添加增量(delta),其中参数delta可以是负数
Done()表示减少WaitGroup计数器的值,应当在程序最后执行.相当于Add(-1)
Wait()表示阻塞直到WaitGroup计数器为0

*/
