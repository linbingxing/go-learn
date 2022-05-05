package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listener err ", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err", err)
		} else {
			fmt.Println("conn success", conn.RemoteAddr())
		}
		process(conn)
	}

}

func process(con net.Conn) {
	defer con.Close()
	for {
		buf := make([]byte, 1024)
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("conn err", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
