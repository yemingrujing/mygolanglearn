package main

import "fmt"

// 可变参数函数
func find(num int, nums ...int) {
	fmt.Printf("type of  nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Println("\n")
}

func change(s ...string) {
	s[0] = "Go"
}

func changes(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func changess(s ...string) {
	s[0] = "Go"
	ss := s[1:2]
	ss = append(ss, "playground")
	fmt.Println(s)
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 209)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}

	// 直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片
	find(89, nums...)

	fmt.Println("=======================================")
	welcome := []string{"hello", "world"}
	// 如果使用了 ..., 切片本身会作为参数直接传入，不需要再创建一个新的切片
	change(welcome...)
	fmt.Println(welcome)

	fmt.Println("=======================================")
	// // 切片引用底层数组，改变切片的值会直接影响原始数组的值，这个前提是切片的数量在cap大小之内，如果切片append元素调整了cap后，是一个新的引用数组，不会改变原来的数组。
	welcomes := []string{"hello", "world"}
	welcomess := []string{"hello", "world", "welcome"}
	changes(welcomes...)
	changess(welcomess...)
	fmt.Println(welcomes)
	fmt.Println(welcomess)
}