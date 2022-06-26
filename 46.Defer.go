package main

import (
	"fmt"
	"os"
)

/*

golang 中的fmt Fprintln实例讲解

在Go语言中，fmt软件包使用与C的printf()和scanf()函数相似的函数来实现格式化的I /O。 Go语言格式的fmt.Fprintln()函数使用其操作数的默认格式并写入w。在此，始终在指定的操作数之间添加空格，并在末尾添加换行符。此外，该函数在fmt包下定义。在这里，您需要导入“fmt”包才能使用这些函数。

用法:

func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

参数：此函数接受两个参数，如下所示：

w,这是指定的标准输入或输出。
a …interface{}:它包含一些代码中使用的字符串和常量变量。
返回值：它返回已写入的字节数以及遇到的任何写入错误。


*/

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {

	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
