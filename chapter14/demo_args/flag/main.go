package main

import(
	"fmt"
	"flag"
)



func main()  {
	var user string
	var pwd string
	var host string
	var port int


	flag.StringVar(&user, "u", "", "用户名输入错误")
	flag.StringVar(&pwd, "pwd", "", "密码输入错误")
	flag.StringVar(&host, "h", "localhost", "主机输入错误")
	flag.IntVar(&port, "p", 8080, "端口输入错误")

	flag.Parse()
	fmt.Println(user, pwd, host, port)
}