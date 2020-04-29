package point_package

import "fmt"

var i int8 = 5

var p *int8 = &i

func PrintPoint() {
	//一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，
	//在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关
	fmt.Printf("i memory address is %p \n", p)
	//你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。
	//当没分配任何值时是nil
	fmt.Printf("value of point is %v \n", *p)
}

func ChangeValue() {
	str := "hello world"
	p := &str
	fmt.Printf("point is %p,vlaue is %v, point value is %v \n", p, str, *p)
	//change value
	*p = "你好，世界"
	fmt.Printf("point is %p,vlaue is %v, point value is %v \n", p, str, *p)
	//你不能得到一个文字或常量的地址 如
	// const i = 5; p = &i 或 p = &10
	// GO语言不允许进行指针运算，可以使用指针进行引用传递
	//多层指针嵌套
	p1 := &p
	fmt.Printf("point is %p, value is %v, value value is %v", p1, *p1, **p1)
	//一个空指针的反向引用是不合法的，并且会使程序崩溃
}
