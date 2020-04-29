package struct_method_package

import "fmt"

//go中的继承是通过匿名结构体实现的

type A struct {
	x1 float64
	y1 float64
}

type B struct {
	A
	x2 float64
	y2 float64
}

func (a *A) abs() {
	if a.x1 < 0 {
		a.x1 = -(a.x1)
	}
	if a.y1 < 0 {
		a.y1 = -a.y1
	}
}

func PrintlnAnonymous() {
	b := B{A{0.1, 0.2}, 0.1, 0.2}
	//b.abs()
	b.A.abs() //如果A是非匿名属性，则无法调用A的方法
	fmt.Printf("b value x1 is %v, y1 is %v, x2 is %v, y2 is %v", b.x1, b.y1, b.x2, b.y2)
}

//外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
//如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。
//没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

//如何在类型上嵌入功能两种方式
//方式一、聚合  包含一个所需功能类型的具名字段。
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log //因为*Log不是匿名属性，所以无法直接调用Log类的方法
}

//如果一个方法返回*Log指针，然后再通过指针调用Log对应的方法
//方式二、内嵌  内嵌（匿名地）所需功能类型 （继承）
/*type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}*/

//很明显的可以看出，内嵌更加有利于对象的重用

//go支持多重继承，通过内嵌多种不同的匿名类型即可

//在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。
