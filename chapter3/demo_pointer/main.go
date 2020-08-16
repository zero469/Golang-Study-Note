package main

import (
	"fmt"
)

func main()  {
	var num int = 1
	var ptr *int = &num
	fmt.Println("address is ",ptr," value is ",*ptr," pointer address is ",&ptr)
	var num2 int = 2
	ptr = &num2
	fmt.Println("address is ",ptr," value is ",*ptr," pointer address is ",&ptr)
	
}