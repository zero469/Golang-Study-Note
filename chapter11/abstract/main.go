package main

import(
	"fmt"
)

//结构体Account
type Account struct{
	AccountNo string
	PassWord string
	Balance float64
}

func (account *Account) checkPwd(pwd string)  bool{
	return account.PassWord == pwd
}
//存款
func (account *Account) Deposite(money float64, pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}

	if money < 0{
		fmt.Println("输入错误")
		return
	}
	account.Balance += money
	
	fmt.Printf("存款%v元成功，当前余额为%v元\n",money, account.Balance)
}

//取款
func (account *Account)Draw(money float64, pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}

	if(money < 0 || money > account.Balance){
		fmt.Println("余额不足！")
		return
	}

	fmt.Printf("取款 %v 元成功\n", money)
	account.Balance -= money
}

//查询余额
func (account *Account)Query(pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}
	fmt.Printf("your balance is %v\n",account.Balance)
}
func inputPassword() (pwd string) {
	fmt.Println("请输入密码：")
	fmt.Scanln(&pwd)
	return 
}

func main()  {
	account := &Account{
		AccountNo : "123456",
		PassWord : "666666",
		Balance : 100,
	}
	var pwd = inputPassword()
	for {
		if !account.checkPwd(pwd){
			pwd = inputPassword()
		}
		fmt.Println("请选择下一步需要执行的操作：")
		fmt.Println("1 : 查询余额")
		fmt.Println("2 : 存款")
		fmt.Println("3 : 取款")
		var opCode int
		var money float64
		fmt.Scanln(&opCode)
		switch opCode {
		case 1:
			account.Query(pwd)
		case 2:
			fmt.Println("请输入存款金额：")
			fmt.Scanln(&money)
			account.Deposite(money, pwd)
		case 3:
			fmt.Println("请输入取款金额：")
			fmt.Scanln(&money)
			account.Draw(money, pwd)
		default:
			fmt.Println("无效指令，请输入正确指令")
			continue
			
		}
	
	}
}