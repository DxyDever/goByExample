package main

import (
	"fmt"
	"time"
)

/*
go请求处理，频率控制
   频率控制是控制资源利用和保证服务高质量的重要机制
   go可以使用goroutine channel ticker来以优雅的方式支持频率控制
*/
func main() {

	/*
		使用一个通道来处理所有的这些请求
	*/
	request := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		request <- i
	}

	close(request)

	/*
		这limiter的Ticker每隔200毫秒结束通道阻塞
		这个limiter就是我们的控制器
	*/
	limiter := time.Tick(200 * time.Millisecond)

	for req := range request {
		//通过这个阻塞，每隔两百毫秒处理一次请求
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()
	/*
		我们可以保持正常的请求频率限制，但是也允许请求短时间内爆发
		我们可以使用通道缓存来实现
		比如下面的这个burstyLimiter就允许同时处理三个事件
	*/
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	/*
		然后每隔两百毫秒再向burstyLimiter发送一个数据，这里不断的每隔200毫秒向
		burstyLimiter发送数据
	*/
	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	/*
		这里模拟的是5个请求burstyRequests的前面三个请求会被连续的进行处理
	*/
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	/*
		前三个请求连续的进行处理，后面的请求会慢慢的，速度会根据这个值来进行变化,time.Tick(2000 * time.Millisecond)
	*/
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
