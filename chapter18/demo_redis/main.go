package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis.Dial err = ",err)
		return 
	}
	defer conn.Close()
	
	fmt.Println("conn succ, conn = ",conn)

	_, err = conn.Do("set", "name", "tom")
	if err != nil{
		fmt.Println("set name err = ",err)
		return
	}
	fmt.Println("set success")
	

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil{
		fmt.Println("set name err = ",err)
		return
		
	}
	fmt.Println(r)
	
}