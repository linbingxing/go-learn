package main

import "fmt"

func sum(n1, n2 int) int {
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)

	n1++
	n2++
	sum := n1 + n2
	fmt.Println("sum", sum)
	return sum
}
func main() {
	fmt.Println(sum(2, 3))
}
