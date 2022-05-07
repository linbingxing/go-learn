package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_learn/src/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("client conn err =", err)
		return
	}

	var msg message.Message

	msg.Type = message.LoginMsgType

	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json marshal err", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json marshal err", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes []byte
	binary.BigEndian.PutUint32(bytes[:4], pkgLen)
	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println(" conn write len err ", err)
		return
	}
	n, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write context err", err)
		return
	}

	return
}
