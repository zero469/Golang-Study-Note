package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)
//全局pool
var pool *redis.Pool

//初始化pool
func init()  {
	pool = &redis.Pool{
		MaxIdle : 8,
		MaxActive : 0,
		IdleTimeout : 100,
		Dial : func() (redis.Conn, error)  {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main()  {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "tom")
	if err != nil{
		fmt.Println("conn.Do err = ",err)
		return
	}
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil{
		fmt.Println("get name err = ",err)
		return
	}
	fmt.Println(r)

	pool.Close()
	conn2 := pool.Get()
	_, err = conn2.Do("set", "name2", "tom2")
	if err != nil{
		fmt.Println("conn.Do err = ",err)
		return
	}
	r, err = redis.String(conn2.Do("get", "name2"))
	if err != nil{
		fmt.Println("get name err = ",err)
		return
	}
	fmt.Println(r)

}