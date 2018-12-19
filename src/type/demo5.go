package main

import "fmt"

/**
复数类型
 */
func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cAdd := c1 + c2
	fmt.Println("sum:", cAdd)
	cMul := c1 * c2
	fmt.Println("product:", cMul)
}
