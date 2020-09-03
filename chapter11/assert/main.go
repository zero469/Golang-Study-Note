package main

import(
	"fmt"
)

type Point struct{
	x int
	y int
}

func main()  {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	fmt.Println(a)

	var b Point
	
	//不能直接转换
	//b = a
	//类型断言
	b = a.(Point)//判断该接口变量是否指向Point类型，如果不是则报错
	fmt.Println(b)

	var x interface{}
	var f float32 = 1.1
	x = f
	var y float32 = x.(float32)//类型需要匹配，否则会panic
	fmt.Printf("y的类型时 %T, 值是 %v\n", y, y)
	fmt.Printf("y的类型时 %T, 值是 %v\n", y, y)



	//断言前进行判断
	var x1 interface{}
	var f1 float64 = 1.2
	x1 = f1
	y1, ok := x1.(float32)
	
	if ok{
		fmt.Println("断言成功")
		fmt.Printf("y1的类型时 %T, 值是 %v\n", y1, y1)
	} else {
		fmt.Println("断言失败")
	}
}