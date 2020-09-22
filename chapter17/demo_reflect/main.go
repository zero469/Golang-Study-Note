package main

import(
	"reflect"
	"fmt"
)


func reflectTest01(a interface{})  {
	
	rTyp := reflect.TypeOf(a)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(a)
	fmt.Println(rVal)
}


func main()  {
	var num int = 100
	reflectTest01(num)
}