package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

/*
端口如何确定，写死在程序里还是系统随机分配？
如果是写死的，那么如果写死的端口被占用怎么办？
还是说系统随机分配，那么分配后怎么通知另一方？
tcp建立连接的过程是什么？
猜测：
客户端发起tcp链接，发起前其端口已被确定，在建立连接的过程中确定两方的端口
*/
func exit()  bool{
	var flag string
	for{
		fmt.Printf("你确定要退出吗？[y/n]")
		fmt.Scanln(&flag)
		switch flag {
		case "y":
			return true
		case "n":
			return false
		default:
			continue
		}
	}
}
func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("client dial err = ", err)
		return
	}
	defer conn.Close()
	fmt.Println("client dial success, conn = ", conn )
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("read string err = ", err)
		}
		//将str发送给服务器
		str = strings.Trim(str, "\n\r")
		// fmt.Printf("%v a", str)
		if str == "exit"{
			if exit(){
				break
			} else{
				continue
			}
		}
		n, err := conn.Write([]byte(str))
		if err != nil{
			fmt.Println("conn.Write err = ", err)
		}
		fmt.Printf("成功发送了%d Byte的数据\n", n)
	}


}