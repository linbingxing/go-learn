package process

import (
	"encoding/json"
	"fmt"
	"go_learn/src/chatroom/client/utils"
	"go_learn/src/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)

	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("conn read context err", err)
		return
	}
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		fmt.Println("登录成功")
		go serverProcessMsg(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return
}
