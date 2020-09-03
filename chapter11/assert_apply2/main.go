package main


import(
	"fmt"
)

type student struct{
	name string
}


//使用类型断言判断参数类型
func typeJudge(items... interface{})  {
	for _, val := range(items){
		switch val.(type){
		case bool:
			fmt.Println("bool: ", val)
			
		case string:
			fmt.Println("string: ", val)
			
		case float32:
			fmt.Println("float32: ", val)
			
		case float64:
			fmt.Println("float64: ", val)
			
		case int, int32, int64:
			fmt.Println("int: ", val)
		
		case student:
			fmt.Println("student: ",val)
		
		case *student:
			fmt.Println("*student: ",val)
		
		default:
			fmt.Println("unclear type: ",val)
		}
	}
}
func main()  {
	typeJudge(1, 1.1, "tom", true, student{"a"}, &student{"b"})
}