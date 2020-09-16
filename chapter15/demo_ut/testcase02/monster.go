package monster

import(
	"encoding/json"
	"io/ioutil"
	"fmt"
)
type monster struct{
	Name string
	Age int
	Skill string
}


func (mons *monster)Store() bool {
	data, err := json.Marshal(mons)
	if err != nil{
		fmt.Printf("序列化错误，err = %v\n", err)
		return false;
	}

	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)//所有人可读可写不可执行
	if err != nil{
		fmt.Printf("写文件错误，err = %v\n", err)
		return false
	}
	return true
}


func (mons *monster)ResStore() bool {
	filePath := "d:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Printf("读文件错误， err = %v\n", err)
		return false
	}

	err = json.Unmarshal(data, mons)
	if err != nil{
		fmt.Printf("反序列话错误，err = %v\n", err)
		return false
	}
	return true
}