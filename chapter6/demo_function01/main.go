package main

import(
	"fmt"
)

var(
	func1 = func (n1 int, n2 int)int  {
		return n1 + n2
	}
)
func main()  {
	//匿名函数
	res := func (n1 int, n2 int) int{
		return n1 + n2
	}(10, 20)
	fmt.Println(res)
	//使用变量保存匿名函数
	a := func(n1 int, n2 int)int{
		return n1 + n2
	}
	fmt.Println(a(1, 2))
	//
	fmt.Println(func1(1, 2))
}