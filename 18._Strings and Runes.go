package main

import (
	"fmt"
	"unicode/utf8"
)

/*
字符串和字符

rune类型字符，代表的是一个utf-8字符，当需要处理中文的时候，就需要使用到rune类型，rune类型
等价于int32

byte类型字符
uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符
golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。


str := "你好a"
//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
fmt.Println("len(str):", len(str))

//以下两种都可以得到str的字符串长度

//golang中的unicode/utf8包提供了用utf-8获取长度的方法
fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

//通过rune类型处理unicode字符

fmt.Println("rune:", len([]rune(str)))

//""   /  ''
//string不可变,不能对string中单独部分进行修改
//such as：string[0]='c'    xxx

//可以string -> byte,修改byte,再转换回string
//可以	s := "hello"
//		s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
//		fmt.Printf("%s\n", s)
//+ connect

//type error
//package errors
//
*/
func main() {

	const s = "你好"

	fmt.Println(len(s))

	fmt.Println()

	fmt.Println("rune count", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Println(runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Println(runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == '你' {
		fmt.Println("found so sua")
	}

}

