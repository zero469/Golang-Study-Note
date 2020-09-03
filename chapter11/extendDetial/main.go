package main

import(
	"fmt"
)

type A struct{
	Name string
	age int
}

func (a *A) SayOk()  {
	fmt.Println("A sayOk", a.Name)
}

func (a *A) hello()  {
	fmt.Println("A say hello", a.Name)
}

type B struct{
	Name string
	A
}


type C struct{
	A 
	B
}
func main()  {
	b := B{}
	//递归向上查找Name属性，首先查找B本身，如果B没有则检查嵌入的A
	//如果B和A有同名属性或方法时，编译器采用就近访问原则，即优先访问B的
	//如果想要访问匿名结构体的属性或方法，那么需要指定即B.A.
	b.Name = "tom"
	b.A.Name = "tom"
	b.age = 19
	b.SayOk()
	b.hello()

	//有多个同名属性时，需要指定属性所在的匿名结构体
	c := C{}
	c.A.Name = "tom"
	c.B.Name = "mary"
	fmt.Println(c)
}