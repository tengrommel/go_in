package main

import (
	"time"
	"fmt"
)

func main() {
	// 首先我们将看一下基本的速率。
	// 假设我们想限制我们接收请求的处理，
	// 我们将这些请求发送给一个相同的通道。
	requests := make(chan int, 5)
	for i:=1;i<=5;i++ {
		requests <-i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)
	// 这个`limite`通道将每200ms接收一个接收，我们限制自己每200ms执行一个请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)
	for i:=0;i<3;i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t:= range time.Tick(time.Millisecond * 200){
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i:=1;i<=5;i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
