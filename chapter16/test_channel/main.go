package main

import(
	"fmt"
)

type Cat struct{
	Name string
	Age int
}

func main()  {
	var allChan chan interface{} = make(chan interface{}, 3)
	allChan<- 10
	allChan<- "string"
	
	allChan<- Cat{"tom", 4}


	<-allChan
	<-allChan

	newCat := <-allChan
	//空接口类型断言
	a , _ := newCat.(Cat)
	fmt.Println(a.Name)
}