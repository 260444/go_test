package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

//newperson构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s正在做梦\n", p.name)
}

//修改年龄的方法	指针类型接收者可以修改
func (p *Person) SetAge(newage int8) {
	p.age = newage
}

//值类型无法修改
func (p Person) SetAge2(newage2 int8) {
	p.age = newage2
}

// 为任意类型添加方法
type myInt int

func (m myInt) SayHello() {
	fmt.Println("int")
}
func main() {
	p1 := NewPerson("小红", 18)
	fmt.Printf("p1.age: %v\n", p1.age)
	p1.Dream()
	p1.SetAge(10)
	fmt.Printf("p1.age: %v\n", p1.age)

	p1.SetAge2(66)
	fmt.Printf("p1.age: %v\n", p1.age)
	//修改失败

	var m3 myInt
	m3 = 123
	m3.SayHello()
	fmt.Printf("%#v %T\n", m3, m3)
}
