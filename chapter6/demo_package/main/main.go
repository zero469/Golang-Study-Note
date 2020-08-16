package main
import(
	"fmt"
	util "go_code/chapter3/demo_package/utils"
)

func main()  {
	fmt.Println(util.Cal(1.2, 1.3, '+'))
	fmt.Println(util.Cal(1.2, 1.3, '-'))
	fmt.Println(util.Cal(1.2, 1.3, '*'))
	fmt.Println(util.Cal(1.2, 1.3, '/'))
	fmt.Println(util.Cal(1.2, 1.3, '='))
}