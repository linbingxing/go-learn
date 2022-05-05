package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile(srcFileName, distFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file error ", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	distFile, err := os.OpenFile(distFileName, os.O_CREATE|os.O_WRONLY, 0666)

	defer distFile.Close()
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	writer := bufio.NewWriter(distFile)
	return io.Copy(writer, reader)
}

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
	//第二种
	context, err := ioutil.ReadFile("D:/vuepress_doc/docs/12.Go学习/Go学习.md")
	if err != nil {
		fmt.Println("file open error", error)
	}
	fmt.Println(string(context))

	file1, err := os.OpenFile("D:/vuepress_doc/docs/12.Go学习/text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file1.Close()
	str := "hello\n"

	writer := bufio.NewWriter(file1)

	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	CopyFile("D:/vuepress_doc/docs/12.Go学习/text.txt", "D:/vuepress_doc/docs/12.Go学习/Go学习.assets/text.txt")
}
