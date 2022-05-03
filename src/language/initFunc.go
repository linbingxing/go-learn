package main

import (
	"fmt"
	"strings"
)

var age = 90

func init() {

	fmt.Println("init", age)
}
func main() {
	fmt.Println("main")

	f := test(".jpg")

	fmt.Println(f("test"))

	fmt.Println(f("t22.jpg"))
}

func test(subffix string) func(name string) string {
	var num = 200
	return func(name string) string {
		num++
		fmt.Println("num", num)
		if !strings.HasSuffix(name, subffix) {
			return name + subffix
		}
		return name
	}
}
