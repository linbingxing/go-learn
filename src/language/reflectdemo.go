package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

/**
反射例子
*/
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value", reflect.ValueOf(x))

	var x1 float64 = 3.4
	v := reflect.ValueOf(x1)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	t := T{23, "test"}

	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("test2")
	fmt.Println("t is now ", t)

}
