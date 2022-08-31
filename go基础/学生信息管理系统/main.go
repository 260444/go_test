package main

import (
	"fmt"
	"os"
)

// 需求：
//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息
func Showmenu() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出")
}

type Student struct {
	ID    int
	Name  string
	Class string
}

// newStudent是student的构造函数
func NewStudent(id int, name, class string) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Class: class,
	}
}

type StudentMgr struct {
	AllStudents []*Student
}

// newStudentmgr是studentmgr的构造函数
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		AllStudents: make([]*Student, 0, 100),
	}
}

//实现添加用户的方法
func (s *StudentMgr) Add(newstu *Student) {
	s.AllStudents = append(s.AllStudents, newstu)
}
func (s *StudentMgr) Modify(newstu *Student) {
	for i, v := range s.AllStudents {
		if newstu.ID == v.ID {
			s.AllStudents[i] = newstu
			return

		}
		fmt.Println("咩有找到", newstu.ID)
	}
}

//实现查询用户的方法
func (s *StudentMgr) Show() {
	for _, v := range s.AllStudents {
		fmt.Printf("学号：%v, 姓名：%v, 班级：%v\n", v.ID, v.Name, v.Class)
	}
}

func GetInput() *Student {
	var (
		id    int
		name  string
		class string
	)

	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员班级:")
	fmt.Scanf("%s ", &class)
	stu := NewStudent(id, name, class)
	return stu
}
func main() {
	sm := NewStudentMgr()
	for {
		// 打印系统菜单
		Showmenu()
		// 等待用户选择要执行的选项
		// 执行用户选择的动作
		var input int
		fmt.Print("请输入你要操作的数字\n")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		switch input {
		case 1:
			s := GetInput()
			sm.Add(s)
			//添加学院
		case 2:
			s := GetInput()
			sm.Modify(s)
			// 编辑学员信息
		case 3:
			// 展示所有学员信息
			sm.Show()
		case 4:
			os.Exit(0)
		}
	}
}
