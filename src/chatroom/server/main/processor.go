package main

import (
	"fmt"
	"go_learn/src/chatroom/common/message"
	process2 "go_learn/src/chatroom/server/process"
	"go_learn/src/chatroom/server/utils"
	"net"
)

type processor struct {
	Conn net.Conn
}

func (pr *processor) Process(conn net.Conn) {
	defer conn.Close()
	for {
		tf := utils.Transfer{
			Conn: conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("conn read err=", err)
			return
		}
		fmt.Println("msg= ", msg)
		pr.ServerProcessMsg(&msg)
	}
}

func (pr *processor) ServerProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		up := process2.UserProcess{
			Conn: pr.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.LoginResMsgType:
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
