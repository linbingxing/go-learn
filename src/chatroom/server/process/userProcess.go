package process

import (
	"encoding/json"
	"fmt"
	"go_learn/src/chatroom/common/message"
	"go_learn/src/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (up *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json unmarshal fail err=", err)
		return
	}
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	var loginResMsg message.LoginResMsg

	if loginMsg.UserId == 100 &&
		loginMsg.UserPwd == "123456" {
		loginResMsg.Code = 200
	} else {
		loginResMsg.Code = 500
		loginResMsg.Error = "用户不存在"
	}
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json marshal fail", err)
		return
	}
	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json marshal fail", err)
		return
	}
	tf := utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}
