package main

import(
	"fmt"
	"encoding/json"
)
type Monster struct{
	Name string `json:"name"`//使用json tag使序列化后的key按tag显示
	Age int `json:"age"`
	Skill string `json:"skill"`
}


func testStruct()  {
	monster := Monster{
		Name : "张三",
		Age : 20,
		Skill : "play",
	}
	//序列化struct
	data, err := json.Marshal(&monster)
	if err != nil{
		fmt.Printf("序列化失败，err=%v\n", err)
	}
	fmt.Println(string(data))
}

func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "张三"
	a["Age"] = 20
	a["Skill"] = "play"

	//序列化map
	data, err := json.Marshal(a)
	if err != nil{
		fmt.Printf("序列化失败，err=%v\n", err)
	}
	fmt.Println(string(data))
	//反序列化
	var um1 map[string]interface{}
	json.Unmarshal(data, &um1)
	fmt.Println(um1)
}

func testSlice()  {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "张三"
	m1["age"] = 21
	m1["skill"] = [...]string{"play", "study"}
 	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "李四"
	m2["age"] = 22
	m2["skill"] = [...]string{"1", "2"}
	slice = append(slice, m2)
	
	data, err := json.Marshal(slice)
	if err != nil{
		fmt.Printf("序列化失败，err=%v\n", err)
	}
	fmt.Println(string(data))
}

//对基本数据类型序列化
func testNormalData()  {
	var num float64 = 3.1415926
	data, err := json.Marshal(&num)
	if err != nil{
		fmt.Printf("序列化失败，err=%v\n", err)
	}
	fmt.Println(string(data))
}

func main()  {
	testStruct()
	testMap()
	testSlice()
	testNormalData()
}