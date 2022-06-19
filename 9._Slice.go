package main

import "fmt"

/*

https://blog.csdn.net/weixin_49411859/article/details/117233947

*/
func main() {

	/*
		切片的长度和容量是不同的概念
	*/
	slice1 := make([]int, 5)

	fmt.Println(slice1, len(slice1), cap(slice1))

	//在原有切片的末尾添加新元素，如果容量不够，就会扩容两倍，前面的元素都会变成默认值
	slice1 = append(slice1, 1)

	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := make([]int, len(slice1))

	copy(slice2, slice1)

	fmt.Println(slice2, len(slice2), cap(slice2))

}

/*
切片的要点需要注意的

- 调用append函数必须用原来的切片变量接受返回值
- 由此可以知道append追加元素，当原来的底层数组放不下的时候，Go底层就会将原来的数组替换，也就是指向一个新的数组中去。
首先判断如果新申请的容量大于2倍的旧容量时，最终容量就是新申请的容量,也就是说不是两倍的旧容量
- 否则判断，如果旧切片长度大于等于1024，则最终容量从旧容量开始循环增加原来的1/4，直到最终容量大于等于新申请的容量。
- 使用copy()函数将切片a1中的元素复制到切片a3,并存储在新的地址
- Go语言中没有删除切片元素的专用方法，因此只能通过自有的方法进行删除
a1 := []int{1,2,3}; // 删除下标为1的数据
a1 = append(a1[:1],a1[2:]...);
fmt.Println(a1)  // [1,3]

a1 := []int{1,2,3}; // 删除下标为1的数据
a2 := a1[:];
fmt.Println(a2,len(a2),cap(a2));  // [1 2 3] 3 3
a2 = append(a1[:1],a1[2:]...);
fmt.Println(a2,len(a2),cap(a2));  // [1 3] 2 3
fmt.Println(a1)  // [1 3 3]

像这样通过切片自身的方法删除的数据会修改顶层数组的数据，原因在于：

切片不保存具体的值
切片对应一个底层数组
底层数组都是占用一块连续的内存
append方法若是返回回来的数组容量不大于原来的数组时，所进行的操作是在同个底层数组中进行操作，同时数组内占用的内存并不会更改，因此数组内被删除数据后的所有数据都左移一位，最后一位数值依旧保留，通俗来讲就是该下标直接覆盖为后面的数据。

var arr = make([]int,5,10);  // 创建时就存在长度为5且值都为0的切片
for i := 0; i < 10; i++{
	arr = append(arr,i);
}
fmt.Println(arr) // [0 0 0 0 0 1 2 3 4 5 6 7 8 9]



*/
