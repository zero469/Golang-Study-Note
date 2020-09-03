package main

import (
	"fmt"
)

//公有结构体
type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名：%v 年龄: %v 成绩: %v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//通过嵌入匿名结构体继承公有结构体的属性和方法
//小学生
type Pupil struct {
	Student
}

func (pupil *Pupil) Testing() {
	fmt.Println("小学生正在考试中....")
}

//大学生
type Graduate struct {
	Student
}
func (graduate *Graduate) Testing() {
	fmt.Println("大学生正在考试中....")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 10
	pupil.Testing()
	pupil.Student.SetScore(90)
	pupil.ShowInfo()
	
	graduate := &Graduate{}
	graduate.Student.Name = "Mary"
	graduate.Student.Age = 20
	graduate.Testing()
	graduate.Student.SetScore(100)
	graduate.Student.ShowInfo()
}
