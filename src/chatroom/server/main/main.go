package main

import (
	"fmt"
	"go_learn/src/chatroom/server/model"
	"net"
)

func initUsrDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	initPool("127.0.0.1:6379", 8, 0, 100)
	initUsrDao()
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	// defer listen.Close()
	for {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn err= ", err)
		}
		pr := processor{
			Conn: conn,
		}
		go pr.Process(conn)

	}
}
