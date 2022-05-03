package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open("D:/vuepress_doc/docs/12.Go学习/Go学习.md")
	if error != nil {
		fmt.Println("file open error", error)
	}
	fmt.Printf("file=%v", file.Name())
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}
