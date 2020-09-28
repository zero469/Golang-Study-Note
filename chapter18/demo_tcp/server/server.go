package main
import(
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	fmt.Printf("服务器等待%v发送消息\n", conn.RemoteAddr().String())
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)//阻塞等待客户发送消息
		if err != nil{
			fmt.Println("Server conn.Read() err = ",err)
			break
		}
		fmt.Println(string(buf[:n]))//打印实际接受到的数据
	}
	fmt.Printf("客户端%v已退出\n", conn.RemoteAddr().String())

}

func main()  {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil{
		fmt.Println("listen error = ", err)
		return
	}
	defer listen.Close()

	//服务器监听
	for{
		fmt.Println("等待连接...")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept() err = ", err)
		}else{
			fmt.Println("Accept() suc con = ", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
	// fmt.Println(listen)
}