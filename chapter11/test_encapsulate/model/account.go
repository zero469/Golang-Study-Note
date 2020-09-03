package model

import (
	"fmt"
)

type account struct {
	accountNo string
	balance   float64
	password  string
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10{
		fmt.Println("账号长度错误")
		return nil
	}
	if len(pwd) != 6{
		fmt.Println("密码长度错误")
		return nil
	}
	if balance < 20{
		fmt.Println("余额错误")
		return nil
	}

	return &account{
		accountNo : accountNo,
		balance : balance,
		password : pwd,
	}
}

//账户的set和get接口
func (account *account) SetAccountNo(no string) {
	if len(no) >= 6 && len(no) <= 10 {
		account.accountNo = no
	} else {
		fmt.Println("账号长度错误")
	}
}

func (account *account) GetAccountNo() string {
	return account.accountNo
}

//密码的set和get接口
func (account *account) SetPassword(pwd string) {
	if len(pwd) == 6 {
		account.password = pwd
	} else {
		fmt.Println("密码长度错误")
	}
}
func (account *account) GetPassword() string {
	return account.password
}

//余额的set和get接口
func (account *account) SetBalance(balance float64) {
	if balance >= 20 {
		account.balance = balance
	} else {
		fmt.Println("余额错误")
	}
}

func (account *account) GetBalance() float64 {
	return account.balance
}


//密码校验
func (account *account) checkPwd(pwd string)  bool{
	return account.password == pwd
}

//存款
func (account *account) Deposite(money float64, pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}

	if money < 0{
		fmt.Println("输入错误")
		return
	}
	account.balance += money
	
	fmt.Printf("存款%v元成功，当前余额为%v元\n",money, account.balance)
}

//取款
func (account *account)Draw(money float64, pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}

	if(money < 0 || money > account.balance){
		fmt.Println("余额不足！")
		return
	}

	fmt.Printf("取款 %v 元成功\n", money)
	account.balance -= money
}

//查询余额
func (account *account)Query(pwd string)  {
	if !account.checkPwd(pwd){
		fmt.Println("密码错误")
		return
	}
	fmt.Printf("your balance is %v\n",account.balance)
}