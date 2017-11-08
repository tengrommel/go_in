package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"runtime"
)

func main() {
	// 我们将使用一个无符号整形数来表示
	var ops uint64 = 0
	for i:=0;i<50;i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
