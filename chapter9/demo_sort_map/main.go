package main

import(
	"fmt"
	// "sort"
)

func main()  {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 15
	map1[8] = 20 
	for k,v := range map1{
		fmt.Println(k, v) 
	}
}