package main

import(
	"fmt"
)
//声明结构体
type cat struct{
	name string
	Age int //首字母大写表示该字段可以在其他包内使用，结构体名同理
	Color string
	Hobby string
	Score [3]int
}


type person struct{
	Name string
	Age int
	Score [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}
func main()  {
	var cat01 cat
	cat01.name = "snow"
	cat01.Age = 1
	cat01.Color = "W"
	cat01.Hobby = "play"
	fmt.Println(cat01)


	//指针 slice map不分配空间初始值是nil
	var p1 person
	fmt.Println(p1)
	fmt.Println(p1.ptr == nil)
	fmt.Println(p1.slice == nil)
	fmt.Println(p1.map1 == nil)

	//使用make分配空间
	p1.slice = make([]int, 10)
	fmt.Println(p1.slice == nil)


}