package main
import (
	"fmt"
	"strconv"
)
func main()  {
	var i int = 100
	fmt.Printf("%T\n",i)
	fmt.Println(strconv.Itoa(i))
	
	
	var str string = "true"
	b , _ :=strconv.ParseBool(str)
	fmt.Println(b)


	var str1 string = "11"
	var c int32
	c, _ = (strconv.ParseInt(str1, 10, 32))
	fmt.Printf("%T, %d\n", c, c)
}
