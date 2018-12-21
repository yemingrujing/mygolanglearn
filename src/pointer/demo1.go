package main

import "fmt"

func change(val *int) {
	*val = 55
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	b := 255
	// 指针变量的类型为 *T,该指针指向一个 T 类型的变量
	// & 操作符用于获取变量的地址.由于 b 可能处于内存的任何位置,你应该会得到一个不同的地址
	a := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	fmt.Println("=======================================")
	// 指针的零值（Zero Value）
	c := 25
	var d *int
	if d == nil {
		fmt.Println("b is", d)
		d = &c
		fmt.Println("b after initialization is", d)
	}

	fmt.Println("=======================================")
	// 指针的解引用:指针的解引用可以获取指针所指向的变量的值,将 a 解引用的语法是 *a
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)

	fmt.Println("=======================================")
	// 向函数传递指针参数
	e := 58
	fmt.Println("value of a before function call is", e)
	f := &e
	change(f)
	fmt.Println("value of a after function call is", e)
	// 不要向函数传递数组的指针，而应该使用切片
	h := [3]int{89, 90, 91}
	modify(h[:])
	fmt.Println(h)
	// Go 不支持指针运算
}