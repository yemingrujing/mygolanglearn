package main

import "fmt"

func main() {
	// 创建了一个名为 personSalary 的 map，其中键是 string 类型，而值是 int 类型
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil, Going to make one.")
		// map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
		personSalary = make(map[string]int)
	}
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	fmt.Println("=======================================")
	// 在声明的时候初始化 map
	personSalary1 := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary1["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	fmt.Println("=======================================")
	employee := "jamie"
	// 获取 map 元素的语法是 map[key]
	fmt.Println("Salary of", employee, "is", personSalary1[employee])

	fmt.Println("=======================================")
	// 如果我们获取一个不存在的元素，会返回 int 类型的零值 0
	fmt.Println("Salary of joe is", personSalary1["joe"])

	fmt.Println("=======================================")
	// value, ok := map[key] 获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在
	newEmp := "joe"
	value, ok := personSalary1[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}

	fmt.Println("=======================================")
	// 遍历 map 中所有的元素需要用 for range 循环
	// 当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同
	for key, value := range personSalary1 {
		fmt.Printf("personSalary1[%s] = %d\n", key, value)
	}
}