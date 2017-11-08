package main

import (
	"time"
	"fmt"
)

func main() {
	// 在我们的例子中，假如我们执行一个外部调用，
	// 并在2秒后通过通道`c1`返回它的执行结果
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	// 这里是使用`select`实现一个超时操作。
	// `res := <- c1`等待结果，`<-Time.After`等待超时
	// 时间1秒后发送的值。由于`select`默认处理第一个已准备好的接收操作，
	// 如果这个操作超过了允许的1秒的话，将会执行超时case。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
