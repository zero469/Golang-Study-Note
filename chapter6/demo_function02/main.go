package main

import(
	"fmt"
	"strings"
)

//闭包
func addUpper() func(int) int{
	var n int = 10
	return func (x int) int{
		n = n + x
		return n
	}
}

func foo1(x *int) func() {
	return func(){
		*x++
		fmt.Printf("foo1 val is %d\n", *x)
	}
}

func foo2(x int) func() {
	return func ()  {
		x++
		fmt.Printf("foo2 val is %d\n", x)	
	}
}


//exec
func makeSuffix(suffix string) func(string) string{
	return func (name string) string {
		if(strings.HasSuffix(name, suffix)){
			return name
		}
		return name + suffix
	}
}
func main()  {
	// f := AddUpper()
	// fmt.Println(f(1))
	// fmt.Println(f(1))
	// fmt.Println(f(1))
	
	// x := 1
	// f1 := foo1(&x)
	// f2 := foo2(x)

	// f1()
	// f2()
	// f1()
	// f2()

	// x = 1

	// f1()
	// f2()
	// f1()
	// f2()

	check := makeSuffix(".jpg")
	fmt.Println(check("a.jp"))
	fmt.Println(check("a"))


}