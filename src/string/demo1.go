package main

import (
	"fmt"
	"unicode/utf8"
)

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

// rune 是 Go 语言的内建类型，它也是 int32 的别称.在 Go 语言中,rune 表示一个代码点.代码点无论占用多少个字节,都可以用一个 rune 来表示
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

// 字符串的长度
func length(s string) {
	// utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度,这个方法传入一个字符串参数然后返回字符串中的 rune 的数量
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

// Go 中的字符串是不可变的,一旦一个字符串被创建，那么它将无法被修改
// 为了修改字符串,可以把字符串转化为一个 rune 切片,然后这个切片可以进行任何想要的改变,然后再转化为一个字符串
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
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

	fmt.Println("=======================================")
	// 字符串的长度
	word1 := "Señor"
	length(word1)
	word2 := "Perts"
	length(word2)

	fmt.Println("=======================================")
	// 字符串是不可变的
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}