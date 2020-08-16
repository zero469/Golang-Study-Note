package main

import(
	"fmt"
)

func deferTest()  {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("ok 1")
}

func deferTest1()  {
	var i = 1
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

func main()  {
	deferTest()
	deferTest1()

}