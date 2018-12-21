package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name string
	salary int
	currency string
	age int
}

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
使用值接收器的方法
 */
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法
 */
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入
 */
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

type address struct {
	city string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName string
	address
}

func main() {
	emp1 := Employee{
		name: "Sam Adolf",
		salary: 5000,
		currency: "$",
	}
	emp1.displaySalary()

	fmt.Println("=======================================")
	// 相同的名字的方法可以定义在不同的类型上,而相同名字的函数是不被允许的
	r := Rectangle{
		length: 10,
		width: 5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius:  12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())

	fmt.Println("=======================================")
	// 指针接收器与值接收器
	// 指针接收器可以使用在:对方法内部的接收器所做的改变应该对调用者可见时; 当拷贝一个结构体的代价过于昂贵时,考虑下一个结构体有很多的字段
	e := Employee {
		name: "Mark Andrew",
		age: 50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)
	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(52)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)

	fmt.Println("=======================================")
	// 匿名字段的方法:属于结构体的匿名字段的方法可以被直接调用,就好像这些方法是属于定义了匿名字段的结构体一样
	p := person{
		firstName: "Elon",
		lastName: "Musk",
		address: address{
			city: "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress()
}