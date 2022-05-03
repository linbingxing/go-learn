package main

import "fmt"

func main() {
	//三种方式
	var m map[string]string
	m = make(map[string]string)
	var a = make(map[int]int, 10)
	for i := 0; i < 20; i++ {
		a[i] = i + i
	}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(m)

	delete(a, 1)
	fmt.Println(a)
	val, ok := a[2]
	if ok {
		fmt.Println(val)
	}

	for k, v := range a {
		fmt.Println(k, v)
	}

	var slice []map[string]string
	slice = make([]map[string]string, 2)
	if slice[0] == nil {
		slice[0] = make(map[string]string, 2)
		slice[0]["name"] = "孙悟空"
		slice[0]["age"] = "500"
	}

	fmt.Println(slice)
}
