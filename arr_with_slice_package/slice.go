package arr_with_slice_package

import "fmt"

//定义：切片（slice）是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），
//所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。
//这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。
//需要注意的是，终止索引标识的项不包括在切片内。切片提供了一个相关数组的动态窗口。
//切片实际上是由三部分组成，指向数组的指针，切片长度以及切片容量

//给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是， 切片的长度可以在运行时修改，
//最小为 0 最大为相关数组的长度：切片是一个 长度可变的数组。
//切片提供了计算容量的函数 cap() 可以测量切片最长可以达到多少：它等于切片从第一个元素开始，到相关数组末尾的元素个数。如果 s 是一个切片
//0 <= len(s) <= cap(s)

//定义：var identifier []type //无需指定长度，一个切片在未初始化之前默认为 nil，长度为 0
var slice1 []int

//初始化 var slice2 = arr2[start:end]
//这表示 slice1 是由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集
var slice2 = arr2[0 : len(arr2)-1]

//上述省略写法
//slice2 = arr2[:]
//另外一种表述 slice1 = &arr2

//其他写法
var slice3 = []int{1, 2, 3}
var slice4 = [...]int{1, 2, 3}
var slice5 = make([]int, 10, 20) //长度是10，容量是20

//切片容量定义：它等于切片从第一个元素开始，到相关数组末尾的元素个数。
// slice2 = slice2[:cap(slice2)]

//说明slice赋值是会赋值整个slice结构体，新的slice与旧的指向同一个数组
func SliceCopy() {
	tempSlice := slice2
	fmt.Printf("tempSlice lenth is %d, cap is %d\n", len(tempSlice), cap(tempSlice))
	tempSlice[0] = 100
	tempSlice[1] = 100
	for idx, val := range slice2 {
		fmt.Printf("slice2 index %d value is %d\n", idx, val)
	}
	for idx, val := range tempSlice {
		fmt.Printf("tempSlice index %d value is %d\n", idx, val)
	}
}

//总结：初始化时：切片长度 = end - start，切片容量 = 目标数组长度 - start
// 重新切片时：如果start =0,则容量不变，否则新的容量 = 旧的容量 - start，新的长度 = end - start
// reslice 过程中start不填默认是当前切片的开头，end不填默认是当前切片的结尾
func Reslice() {
	fmt.Printf("slice2 length is %d cap is %d \n", len(slice2), cap(slice2))
	slice2 = slice2[1:] //长度减1,表示容量减1 因为新的start变成1
	fmt.Printf("slice2 length is %d cap is %d \n", len(slice2), cap(slice2))
	slice2 = slice2[:cap(slice2)] //容量不变，长度变成与容量一致
	fmt.Printf("slice2 length is %d cap is %d \n", len(slice2), cap(slice2))
	slice2 = slice2[:1] //长度变成1，容量不变
	fmt.Printf("slice2 length is %d cap is %d \n", len(slice2), cap(slice2))
	slice2 = slice2[0:] //容量不变，长度也不变
	fmt.Printf("slice2 length is %d cap is %d \n", len(slice2), cap(slice2))
}

//垃圾回收
//切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。
