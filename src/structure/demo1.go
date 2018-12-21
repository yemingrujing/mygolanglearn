package main

import "fmt"
import "structure/computer"

// 创建命名的结构体
type Employee struct {
	firstName, lastName string
	age, salary int
}

// 匿名字段
type NPerson struct {
	string
	int
}

type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	address Address
}

type Persons struct {
	name string
	age int
	Address
}

type name struct {
	firstName, lastName string
}

func main() {

	emp1 := Employee{
		firstName: "Sam",
		age: 25,
		salary: 500,
		lastName: "Anderson",
	}
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	fmt.Println("=======================================")
	// 创建匿名结构体
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	fmt.Println("=======================================")
	// 结构体的零值（Zero Value）
	var emp4 Employee
	fmt.Println("Employee 4", emp4)
	emp5 := Employee{
		firstName: "John",
		lastName: "Paul",
	}
	fmt.Println("Employee 5", emp5)

	fmt.Println("=======================================")
	// 访问结构体的字段
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	fmt.Println("=======================================")
	// 结构体的指针
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	fmt.Println("First Name", emp8.firstName)
	fmt.Println("Age", emp8.age)

	fmt.Println("=======================================")
	// 匿名字段:当我们创建结构体时,字段可以只有类型,而没有字段名.这样的字段称为匿名字段（Anonymous Field）
	// 虽然匿名字段没有名称,但其实匿名字段的名称就默认为它的类型
	np := NPerson{"Naveen", 50}
	fmt.Println(np)
	var np1 NPerson
	np1.string = "naveen"
	np1.int = 50
	fmt.Println(np1)

	fmt.Println("=======================================")
	// 嵌套结构体（Nested Structs）
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name", p.name)
	fmt.Println("Age", p.age)
	fmt.Println("City", p.address.city)
	fmt.Println("State", p.address.state)

	fmt.Println("=======================================")
	// 提升字段（Promoted Fields）:如果是结构体中有匿名的结构体类型字段,则该匿名结构体里的字段就称为提升字段
	var ps Persons
	ps.name = "Naveen"
	ps.age = 50
	ps.Address = Address{
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name", ps.name)
	fmt.Println("Age", ps.age)
	fmt.Println("City", ps.city)
	fmt.Println("State", ps.state)

	fmt.Println("=======================================")
	// 导出结构体和字段:如果结构体名称以大写字母开头,则它是其他包可以访问的导出类型（Exported Type）.同样,如果结构体里的字段首字母大写,它也能被其他包访问到.
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)

	fmt.Println("=======================================")
	// 结构体相等性（Structs Equality）:结构体是值类型。如果它的每一个字段都是可比较的,则该结构体也是可比较的.如果两个结构体变量的对应字段相等,则这两个变量也是相等的
	// 如果结构体包含不可比较的字段,则结构体变量也不可比较
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}