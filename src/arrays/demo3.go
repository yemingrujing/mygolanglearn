package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

// 当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见
func main() {
	// creates and array and returns a slice reference
	nos := []int{9, 8, 7, 6, 5}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)

	// 类似于数组，切片可以有多个维度
	fmt.Println("=======================================")
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Println("\n")
	}
}