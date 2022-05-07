package main

import (
	"fmt"
	"go_learn/src/chatroom/client/process"
	"os"
)

var userId int
var userPwd string

func main() {
	var key int

	var loop = true

	for loop {
		fmt.Println("----------欢迎登录聊天室-------------")
		fmt.Println("\t\t\t1、登录聊天室")
		fmt.Println("\t\t\t2、注册")
		fmt.Println("\t\t\t3、退出系统")
		fmt.Println("\t\t\t请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册")
		case 3:
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
