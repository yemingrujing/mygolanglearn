package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function:", num)
}

func main() {
	var a[3] int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
	fmt.Println("=======================================")
	b := [3]int{12, 78, 50}
	fmt.Println(b)
	fmt.Println("=======================================")
	c := [3]int{12}
	fmt.Println(c)
	fmt.Println("=======================================")
	d := [...]int{12, 78, 50}
	fmt.Println(d)
	fmt.Println("=======================================")
	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e
	f[0] = "Singapore"
	fmt.Println("e is:", e)
	fmt.Println("f is:", f)
	fmt.Println("=======================================")
	num := [...]int{5, 6, 7, 8, 9}
	fmt.Println("before passing to function:", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)
	fmt.Println("=======================================")
	g := [...]float64{67.7, 89.9, 21, 78, 56}
	fmt.Println("length of g is:", len(g))
	for i :=0; i < len(g); i++ {
		fmt.Printf("%d the element of g is %.2f\n", i, g[i])
	}
	sum := float64(0)
	// 另外一种循环数组的方法
	fmt.Println("=======================================")
	for i, v := range g {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
