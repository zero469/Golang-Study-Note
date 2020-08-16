package main

import(
	"fmt"
)

func main()  {
	fmt.Println("ok1")
	goto x
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	x: 
	fmt.Println("ok6")
	fmt.Println("ok7")
	return 
	fmt.Println("ok8")
}