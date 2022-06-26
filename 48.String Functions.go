package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	//strings.Join用于元素类型为string的slice使用分割符号拼接组成一个字符串
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	//这里的第三个参数是替换次数，如果n小于0，则替换次数没有限制
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	//分割相对来说比较简单，go语言的strings包中封装了分割函数strings.Split函数返回的是一个切片slice
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
