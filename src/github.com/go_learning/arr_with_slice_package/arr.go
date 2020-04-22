package arr_with_slice_package

import (
	"fmt"
	"reflect"
)

//数组
//数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列（这是一种同构的数据结构）；
//这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。
//数组长度必须是一个常量表达式，并且必须是一个非负整数。
//数组长度也是数组类型的一部分，所以 [5] int 和 [10] int 是属于不同类型的
//数组最大长度是2Gb
//声明格式：var identifier [len]type

//初始化方式
var arr1 [5]int //初始化0值
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = []int{1, 2, 3, 4, 5}    //第0位是1，第一位是2....
var arr4 = []int{5: 1, 4: 2, 3: 3} //长度是5，注意值是从后往前赋值的，第四位是1，第三位2
var arr5 = new([5]int)             //创建数组指针

//遍历
func InterArr() {
	fmt.Printf("arr1 type is:%v \n", reflect.TypeOf(arr1))
	for idx, val := range arr3 {
		fmt.Printf("arr3的第%d位置值是:%d\n", idx, val)
	}
	for idx, val := range arr4 {
		fmt.Printf("arr4的第%d位置值是：%d\n", idx, val)
	}
	for idx, val := range arr5 {
		fmt.Printf("arr5的第%d位置值是：%d\n", idx, val)
	}
}

//数组拷贝，这个列子我们可以看出数组拷贝是值拷贝，更改新的数组内容不会影响旧的数组
func CopyArr1() {
	tempArr1 := arr1
	fmt.Printf("old arr1 value is %d, old tempArr1 value is %d \n", arr1[0], tempArr1[0])
	tempArr1[0] = 10
	fmt.Printf("new arr1 value is %d, new tempArr1 value is %d \n", arr1[0], tempArr1[0])
}

//切片拷贝会影响原先的数组值，arr3,arr4其实技术实现上已经是切片而非数组
func CopyArr2() {
	tempArr3 := arr3
	fmt.Printf("old arr3 value is %d, old tempArr3 value is %d \n", arr3[0], tempArr3[0])
	tempArr3[0] = 10
	fmt.Printf("new arr3 value is %d, new tempArr3 value is %d \n", arr3[0], tempArr3[0])
}

//指针数组拷贝会影响原先数组，但是不会拷贝整个数组内容,而只是拷贝指针
func CopyArr3() {
	tempArr5 := arr5
	//arr5[0] 由于arr5其实是一个数组指针，go编译器会自动判断你是取指针还是取值，其实arr5[0]在编译器内部是做了处理的
	fmt.Printf("old arr5 value is %d, old tempArr5 value is %d \n", arr5[0], tempArr5[0])
	tempArr5[0] = 10
	fmt.Printf("new arr5 value is %d, new tempArr5 value is %d \n", arr5[0], tempArr5[0])
}

//取别名
type Vector3D [3]float64

var vec Vector3D

//多维数组
const WIDTH = 100
const HEIGHT = 500

var screen = [WIDTH][HEIGHT]int{
	0: {1, 2, 3, 4},
}

//传递大数组作为方法入参的两种方式
//1.传递数组的指针
//2.使用数组的切片
