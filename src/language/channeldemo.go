package main

import "fmt"

func main() {
	var chanint chan int
	chanint = make(chan int, 10)
	for i := 0; i < 10; i++ {
		chanint <- i
	}
	//close(chanint)
	//for v :=  range chanint{
	//	fmt.Println("v=",v)
	//}

	for {
		select {
		case v := <-chanint:
			fmt.Println("v=", v)
		default:
			fmt.Println("取不了")
			return
		}
	}
}
