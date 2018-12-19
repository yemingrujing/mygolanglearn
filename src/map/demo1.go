package main

import "fmt"

func compareMap(m1, m2 map[string]int)(result bool) {
	if len(m1) != len(m2) {
		result = false
		return
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if ok == true && v2 == v1 {
			continue
		} else {
			result = false
			return
		}
	}
	return true
}

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

	fmt.Println("=======================================")
	// 删除 map 中的元素
	fmt.Println("map before deletion", personSalary1)
	delete(personSalary1, "steve")
	fmt.Println("map after deletion", personSalary1)

	fmt.Println("=======================================")
	// 获取 map 的长度
	personSalary1["steve"] = 12000
	fmt.Println("length is", len(personSalary1))

	fmt.Println("=======================================")
	// map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量
	fmt.Println("Original person salary", personSalary1)
	newPersionSalary := personSalary1
	newPersionSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary1)

	fmt.Println("=======================================")
	// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil, 判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素
	map1 := map[string]int {
		"one": 1,
		"two": 2,
	}
	map2 := map1
	result := compareMap(map1, map2)
	fmt.Println("map1 compare map2:", result)

	map3 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	result = compareMap(map1, map3)
	fmt.Println("map1 compare map3:", result)
}