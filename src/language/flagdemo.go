package main

import (
	"flag"
	"fmt"
)

//获取命令行参数
func main() {

	var user string
	flag.StringVar(&user, "u", "", "user")
	flag.Parse()
	fmt.Println(user)
	fmt.Println("222222222222223")
}
