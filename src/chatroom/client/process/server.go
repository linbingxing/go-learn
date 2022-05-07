package process

import (
	"fmt"
	"go_learn/src/chatroom/client/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("---------登录成功--------")
	fmt.Println("---------1、显示在线用户列表--------")
	fmt.Println("---------2、发送信息--------")
	fmt.Println("---------3、信息列表--------")
	fmt.Println("---------4、退出系统--------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误")

	}
}

func serverProcessMsg(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端读取消息错误", err)
			return
		}
		fmt.Printf("接收消息%v", msg)
	}

}
