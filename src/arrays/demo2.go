package main

import "fmt"

func main() {
	a := [5]int{76, 77, 78, 78, 80}
	var b = a[1:4]
	fmt.Println(b)

	fmt.Println("=======================================")
	i := make([]int, 5, 5)
	fmt.Println(i)

	// 追加切片元素
	fmt.Println("=======================================")
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	//切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片
	fmt.Println("=======================================")
	var names []string
	if names == nil{
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// 使用 ... 运算符将一个切片添加到另一个切片
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	fmt.Println("=======================================")
	c := []int{1, 2, 3}
	// Capacity 偶数增加
	c = append(c, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c))
}