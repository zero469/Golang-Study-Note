package main

import(
	"fmt"
)

func main()  {
	// for i := 1; i < 10; i++{//i作用域在循环内
	// 	fmt.Println(i)
	// }
	// i := 1
	// fmt.Println(i)
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// for{
	// 	fmt.Println("in loop")
	// 	break;
	// }
	// var str = "hello world!北京"
	// for i = 0; i < len(str); i++{//按字节遍历
	// 	fmt.Printf("%c ",str[i])
	// 	fmt.Println(unsafe.Sizeof(str[i]))
	// }
	// var newStr = []rune(str)//转成rune切片
	// for i = 0; i < len(newStr); i++{//按字节遍历
	// 	fmt.Printf("%c ",newStr[i])
	// 	fmt.Println(unsafe.Sizeof(newStr[i]))
	// }

	// //for range遍历，按字符遍历
	// for index, val := range str{
	// 	fmt.Printf("%d %c ", index, val)
	// 	fmt.Println(unsafe.Sizeof(val))
	// }
	
	//打印金字塔
	/*
	*
	**
	***
	*/
	var depth int
	fmt.Println("请输入层数")
	fmt.Scanln(&depth)
	for i := 1; i <= depth; i++{
		for j := 1; j <= i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	/*
	  *
	 ***
	*****
	*/
	for i := 1; i <= depth; i++{
		for j := 1; j <= depth - i; j++{
			fmt.Printf(" ")
		}
		for j := 1; j <= 2 * i - 1; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	/*
	  *
	 * *
	*****
	*/
	for i := 1; i <= depth; i++{
		for j := 1; j <= depth - i; j++{
			fmt.Printf(" ")
		}
		for j := 1; j <= 2 * i - 1; j++{
			if(j == 2 * i - 1 || j == 1 || i == depth){
				fmt.Printf("*")
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	

}