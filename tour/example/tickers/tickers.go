package main

import (
	"time"
	"fmt"
)

func main() {
	// 打点器和定时器的机制有点相似：一个通道用来发送数据。
	// 这里我们在这个通道上使用内置的`range`来迭代值每隔500ms发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t:= range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
