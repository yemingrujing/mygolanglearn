package main

import "fmt"

// 单独获取字符串的每一个字节
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

// rune 是 Go 语言的内建类型，它也是 int32 的别称
func printChar(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}

// 字符串的 for range 循环
func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}


func main() {
	name := "Hello World"
	fmt.Println(name)

	fmt.Println("=======================================")
	printBytes(name)
	fmt.Println("\n")

	fmt.Println("=======================================")
	printChars(name)
	fmt.Println("\n")

	fmt.Println("=======================================")
	name = "Señor"
	printBytes(name)
	fmt.Println("\n")
	printChars(name)
	fmt.Println("\n")
	printChar(name)
	fmt.Println("\n")

	fmt.Println("=======================================")
	printCharsAndBytes(name)
	fmt.Println("\n")

	fmt.Println("=======================================")
	// 用字节切片构造字符串
	byteSlice1 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str1 := string(byteSlice1)
	fmt.Println(str1)
	// decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	byteSlice2 := []byte{67, 97, 102, 195, 169}
	str2 := string(byteSlice2)
	fmt.Println(str2)

	fmt.Println("=======================================")
	// 用 rune 切片构造字符串
	runeSlice3 := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str3 := string(runeSlice3)
	fmt.Println(str3)
}