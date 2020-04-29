package interface_package

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

//接口提供了一种方式来 说明 对象的行为：如果谁能搞定这件事，它就可以用在这儿。
//接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。
//接口里也不能包含变量。

//（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，
//例如 Printer、Reader、Writer、Logger、Converter 等等。
//还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开

//类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
//实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
//一个类型可以实现多个接口。
//接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

func Reader() {
	var r io.Reader
	r = os.Stdin // see 12.1
	r = bufio.NewReader(r)

	r = new(bytes.Buffer)

	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
}

type ValueAble interface {
	setValue(arg int)
	getValue() (result int)
}

//可以有多个一样的接口
/*type ValueAble2 interface {
	setValue(arg int)
	getValue()(result int)
}*/

type A struct {
	properties1 int
}

func (a *A) setValue(arg int) {
	a.properties1 = arg
}

func (a *A) getValue() (result int) {
	result = a.properties1
	return
}

//B 可以覆盖A的接口实现
/*func (a *B)setValue(arg int)  {
	a.properties1 = arg * 10
}

func (a *B)getValue() (result int) {
	result = a.properties1 *10
	return
}*/

type B struct {
	A
	properties2 int
}

func TestInterface() {
	a := A{1}
	//不能直接赋值给ValueAble，
	//需要赋值指针给ValueAble,因为实现方法的接受者是指针类型，不会自动生成值类型的方法
	//如果是值类型的则会自动生成指针类型的方法
	var va ValueAble = &a
	fmt.Printf("a vlaue is %v\n", va.getValue())
	a.setValue(10)
	fmt.Printf("a vlaue is %v\n", va.getValue())

	b := B{A: A{properties1: 2}, properties2: 2}
	var vb ValueAble = &b
	fmt.Printf("b vlaue is %v\n", vb.getValue())
	b.setValue(20)
	fmt.Printf("b vlaue is %v\n", vb.getValue())

}

//接口内嵌接口
//一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

//接口断言，判断某个类型是否实现了某个接口
type list []int

func (l list) getValue() (result int) {
	return l[0]
}

func (l list) setValue(i int) {
	l[0] = i
}

func InterfaceAsset() {
	//TODO 搞清楚两者区别
	//结构体不能直接赋值给接口，需要赋值指针给接口，而数组可以直接赋值给接口
	/*l := list([]int{1,2})
	var v ValueAble = l*/
	a := new(A)
	var va ValueAble = a
	if t, ok := va.(*A); ok {
		//输出interface_package.*A
		fmt.Printf("va implement type is %T\n", t)
	}

	b := new(B)
	var vb ValueAble = b
	if t, ok := vb.(*B); ok {
		fmt.Printf("vb implement type is %T\n", t)
	}
}

//类型判断
func TypeSwitch() {
	a := new(A)
	var va ValueAble = a
	switch t := va.(type) {
	case *A:
		fmt.Printf("va type is *A\n")
	case *B:
		fmt.Printf("va type is *B\n")
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
