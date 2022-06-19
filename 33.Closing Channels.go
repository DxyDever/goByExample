package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {

		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)

	fmt.Println("send all jobs")

	//只要往done里面写了东西，立马就会放行
	<-done
}

/*

不指定长度的channel，是非缓冲阻塞管道
指定长度的channel，是缓冲非阻塞管道

下面来看一个例子

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
		close(ch)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("ch length is ", len(ch))
	<-ch
}

是 1 吗？ 答案是： 0， 为什么呢？

我们来分析下，首先在主协程中创建了一个通道，然后在子协程中往这个通道中发送内容，发送完毕后关闭通道，然后在主协程中等待 5s, 等待的目的是让子协程能够运行起来，随后我们打印通道的长度。

我们知道对于非缓冲通道而言，发送方和接收方必须同时准备好，这样数据才能发送过去，而在上面代码中我们先打印通道的长度，然后主协程作为接收方才会从通道中取消息内容，所以我们打印通道的时候由于接收方还没有准备接收，从而导致发送方一直处于阻塞状态，此时消息是没有进入到通道的，所以此时的通道长度为 0 。

非缓冲通道并不需要内存空间来存储发送的消息内容，它在发送方和接受方都准备好时，直接将发送方的消息内容拷贝到接收方。

缓冲通道在容量够的时候，写入操作并不会阻塞，是否阻塞强依赖于读取操作，发送方发送数据后直接返回，这样就决定了缓冲通道需要一块存储空间来真正存储要发送的消息内容。

在多核 CPU 下的 Go 多协程环境，必然会出现多个协程同时并行的对同一个 channel 进行读取操作的情况，这样就带来了资源并发访问的问题。

对于非缓冲通道，发送方与接收方必须是一对一、点对点的操作。如果有多个协程同时对某一个非缓冲 channel 进行操作，那么就会发生有部分读/写操作不能匹配到相应的写/读操作，必然会导致死锁。

对于缓冲通道，由于 Go 采用对通道接收的锁机制，即在接收方在取通道内容时，先用锁把通道锁住，然后取数据，取完数据之后再释放锁，这样就保证了多个协程之间是顺序执行的，不会发生多个协程获取通道里的同一个数据的问题，在某一个协程进行接收操作的时候，其他协程只能阻塞在那里，等待上一个协程释放锁，这样就保证了通道数据的并发安全性。

同样发送方也遵循这样的锁机制，即保证同一时刻只有一个协程来操作通道。

*/
