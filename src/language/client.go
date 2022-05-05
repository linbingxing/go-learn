package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	fmt.Println("conn success ", conn.RemoteAddr())

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("reader error ", err)
		}
		line = strings.Trim(line, " \t\n")
		if line == "exit" {
			break
		}
		conn.Write([]byte(line + "\n"))
	}

}
