package main

import(
	"fmt"
)
//接口声明
type Usb interface{
	//声明没有实现的方法
	Start()
	Stop()
}


type Phone struct{

}
//实现USB接口的所有方法
func (phone *Phone)Start()  {
	fmt.Println("手机开始工作...")
}
func (phone *Phone)Stop()  {
	fmt.Println("手机停止工作...")
}


type Camera struct{

}
//实现USB接口的所有方法
func (camera *Camera)Start()  {
	fmt.Println("相机开始工作...")
}
func (camera *Camera)Stop()  {
	fmt.Println("相机停止工作...")
}



//计算机
type Computer struct{

}
func (c *Computer) Working(usb Usb)  {
	usb.Start()
	usb.Stop()
}

func main()  {
	p := &Phone{}
	c := &Camera{}
	com := &Computer{}	
	com.Working(p)
	com.Working(c)
}