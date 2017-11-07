package main

import "fmt"

func f(from string)  {
	for i:=0;i<3;i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 假设我们有一个函数叫做`f(s)`。我们使用一般的方式
	f("direct")
	// 使用 `go f(s)`在一个Go协程中调用这个函数
	// 这个新的Go协程
}
