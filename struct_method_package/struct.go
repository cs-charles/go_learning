package struct_method_package

import (
	"fmt"
	"unsafe"
)

/**
格式定义
构体里的字段都有 名字，像 field1、field2 等，如果字段在代码中从来也不会被用到，那么可以命名它为 _
type identifier struct {
    field1 type1
    field2 type2
    ...
}
*/

type People struct {
	name string
	age  uint8
}

var p0 People
var p1 *People = new(People)
var p2 People = People{name: "xy", age: 18}
var p3 *People = &People{name: "xy", age: 18}

//选择器：不管是结构体类型还是结构体类型指针，都可以通过符号 '.'进行运算

func InitStructValue() {
	fmt.Printf("struct people init value name is %s,age is %d\n", p0.name, p0.age)
}

//Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。不像 Java 中的引用类型，
//一个对象和它里面包含的对象可能会在不同的内存空间中，这点和 Go 语言中的指针很像。下面的例子清晰地说明了这些情况：
//type Rect1 struct {Min, Max Point }
//type Rect2 struct {Min, Max *Point }

//结构体转换
//机构体别名需要强行的类型转换
func Converter() {
	type p People
	p4 := p{name: "hh", age: 12}
	var p5 People = People(p4)
	var p6 p = p(p5)
	//变量处于不同的内存地址
	fmt.Printf("memomery site p4 is %x,p5 is %x, p6 is %v\n", unsafe.Pointer(&p4), unsafe.Pointer(&p5), unsafe.Pointer(&p6))
	fmt.Printf("struct people init value name is %s,age is %d\n", p5.name, p5.age)
	fmt.Printf("struct people init value name is %s,age is %d\n", p6.name, p6.age)

}

func Compare() {
	p4 := p2
	p5 := p3
	fmt.Printf("struct p2 people init value name is %s,age is %d\n", p2.name, p2.age)
	fmt.Printf("struct p4 people init value name is %s,age is %d\n", p4.name, p4.age)
	fmt.Printf("struct p3 people init value name is %s,age is %d\n", p3.name, p3.age)
	fmt.Printf("struct p5 people init value name is %s,age is %d\n", p5.name, p5.age)

	//p4的值改变不会引起p2改变，因为p4是值类型
	p4.name = "xy1"
	//p5的值改变会引起p3的值改变
	p5.name = "xy1"
	fmt.Printf("struct p2 people init value name is %s,age is %d\n", p2.name, p2.age)
	fmt.Printf("struct p4 people init value name is %s,age is %d\n", p4.name, p4.age)
	fmt.Printf("struct p3 people init value name is %s,age is %d\n", p3.name, p3.age)
	fmt.Printf("struct p5 people init value name is %s,age is %d\n", p5.name, p5.age)

}
