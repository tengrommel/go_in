package main
import "fmt"

type rect struct  {
	width, height int
}

// 这里的 `area`方法有一个_接收器类型_ `rect`

func (r *rect)area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法，这里是一个
// 值类型接收器的例子。
func (r rect) perim() int {
	return 2 * r.width + 2*r.height
}

func main() {
	r := rect{width:10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
