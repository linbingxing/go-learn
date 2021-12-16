package main

import (
	"fmt"
	"sync"
)

func main() {

	// 第一种, var + 变量名 + 变量类型
	var a int = 3

	// 第二种, var + 变量名
	var b = 3.1415926

	// 第三种, 变量名 := 值
	c := "Hello"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 第四种，定义多个变量
	var a1, b1, c1 = 12, "go", false

	//简化初始化
	a2, b2, c2 := 123, "golang", true

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	fmt.Println("变量初始化:")
	//变量初始化
	var a3 int         // a的值为: 0      类型为: int
	var b3 string      // b的值为: ""     类型为: string
	var c3 interface{} // c的值为: nil    类型为: interface{}
	var d3 sync.Pool   // d的值为: Pool{} 类型为: sync.Pool
	var e3 *sync.Pool  // e的值为: nil    类型为: *sync.Pool

	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(c3)
	fmt.Println(d3)
	fmt.Println(e3)
}
