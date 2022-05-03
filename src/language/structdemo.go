package main

import "fmt"

type student struct {
	Name    string
	Age     int
	Address string
}

func stud(stu *student) {
	stu.Age = 600
}
func (stu *student) talk() {
	stu.Age = 1000
}
func main() {
	var stu student
	stu.Name = "孙悟空"
	stu.Age = 500
	stu.Address = "花果山"
	stud(&stu)
	fmt.Println(stu)
	var stu1 *student = new(student)
	stu1.Name = "孙悟空"
	fmt.Println(stu1)
	var stu2 *student = &student{"孙悟空", 500, ""}
	stu2.Name = "孙悟空"
	(*stu2).Name = "孙悟空1"
	fmt.Println(stu2)
	var stu3 student = student{
		"孙悟空",
		500,
		"花果山",
	}
	fmt.Println(stu3)
	stu3.talk()
	fmt.Println(stu3)
}
