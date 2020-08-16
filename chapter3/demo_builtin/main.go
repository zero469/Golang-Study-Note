package main

import(
	"fmt"
)

func main()  {
	num := new(int)
	fmt.Printf("%T, %v, %v %v\n", num, num, *num, &num)
}