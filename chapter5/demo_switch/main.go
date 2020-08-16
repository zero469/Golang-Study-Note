package main

import "fmt"

func main()  {
	// var a = 6
	// switch a{
	// 	case 10:
	// 		//不需要break，同一个case内可以有多个表达式
	// 		fmt.Println("in case 10")
	// 	case 9, 8, 7, 6:
	// 		fmt.Println("in case 9, 8, 7, 6")
	// 	default:
	// 		fmt.Println("in case default")
	// }
	var char byte
	fmt.Scanf("%c", &char)
	switch char{
	case 'a', 'b', 'c', 'd':
		fmt.Println(string(char + 'A' - 'a'))
	default:
		fmt.Println("other")
	}
}