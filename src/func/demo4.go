package main

import "fmt"

/**
空白符：_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
 */
func rectProps2(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, _ := rectProps2(10.8, 5.6)
	fmt.Printf("Area is %f", area)
}
