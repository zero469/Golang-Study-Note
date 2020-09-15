package main

import(
	"os"
	"fmt"
)
//go build main.go
//main.exe [args]

func main()  {

	for _, val := range(os.Args){
		fmt.Println(val)
	}
}