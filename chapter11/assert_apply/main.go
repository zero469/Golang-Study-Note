package main

import(
	"fmt"
)
//Usb 接口声明
type Usb interface{
	//声明没有实现的方法
	Start()
	Stop()
}

//Phone 结构体
type Phone struct{
	name string
}
//Start 实现USB接口的Start方法
func (phone Phone)Start()  {
	fmt.Println("手机开始工作...")
}
//Stop 实现USB接口的Stop方法
func (phone Phone)Stop()  {
	fmt.Println("手机停止工作...")
}
//Call 方法
func (phone Phone)Call()  {
	fmt.Println("通话中...")
}

//Camera 结构体
type Camera struct{
	name string
}
//Start 实现USB接口的Start方法
func (camera Camera)Start()  {
	fmt.Println("相机开始工作...")
}
//Stop 实现USB接口的Stop方法
func (camera Camera)Stop()  {
	fmt.Println("相机停止工作...")
}



//Computer 计算机
type Computer struct{

}
//Working 函数可以接受实现Usb接口的所有实参类型
func (c Computer) Working(usb Usb)  {
	usb.Start()
	phone, ok := usb.(Phone)
	if ok{
		phone.Call()
	} 
	usb.Stop()
}

func main()  {
	//多态数组，使用Usb数组存放Phone和Camera变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"huawei"}
	usbArr[1] = Phone{"apple"}
	usbArr[2] = Camera{"nikon"}
	
	//遍历usb数组，当遍历到Phone类型时调用Call方法
	var computer Computer
	for _, v := range(usbArr){
		computer.Working(v)
	}
	fmt.Println(usbArr)

}