package main

import(
	"fmt"
)

type Ainterface interface{
	Test01()
	Test02()
}

type Binterface interface{
	Test01()
	Test03()
}

type Cinterface interface{
	Ainterface
	Binterface
}
//实现Cinterface
type Stu struct{

}
func (s Stu)Test01()  {
	fmt.Println("Test01")
}

func (s Stu)Test02()  {
	fmt.Println("Test02")
}
func (s Stu)Test03()  {
	fmt.Println("Test03")
}


func main()  {
	stu := Stu{}
	var c Cinterface = stu
	var b Binterface = stu
	var a Ainterface = stu
	c.Test01()
	c.Test02()
	c.Test03()

	b.Test01()
	b.Test03()

	a.Test01()
	a.Test02()
	fmt.Println("hello")
}
