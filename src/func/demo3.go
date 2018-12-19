package main

import "fmt"

/**
命名返回值：从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
 */
func rectProps1(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}

func main() {
	area, perimeter := rectProps1(10.8, 5.6)
	fmt.Printf("Area is %f, Perimeter is %f", area, perimeter)
}
