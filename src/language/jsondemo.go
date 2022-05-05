package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func testStruct() {
	stu := Stu{
		Name: "孙悟空",
		Age:  500,
	}
	data, err := json.Marshal(&stu)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(data))
}

func testMap() {
	m := make(map[string]string, 10)
	m["key1"] = "v1"
	m["key2"] = "v2"
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(data))
	t := make(map[string]string, 10)
	err1 := json.Unmarshal(data, &t)
	if err1 != nil {
		fmt.Println(err1)
	}
	for k, v := range t {
		fmt.Println(k, v)
	}
}

func testSlice() {
	//var slice []map[string]string

}
func main() {
	testStruct()
	testMap()
}
