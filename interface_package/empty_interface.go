package interface_package

import "fmt"

//空接口或者最小接口 不包含任何方法，它对实现不做任何要求：
type Any interface{}

//空接口类似 Java/C# 中所有类的基类： Object 类，二者的目标也很相近。
//
//可以给一个空接口类型的变量 var val interface {} 赋任何类型的值。

func TestEmptyInterface() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = "哈哈哈哈哈"
	fmt.Printf("val has the value: %v\n", val)
	a := new(A)
	a.properties1 = 2
	val = a
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *A:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

//构建一个含有不同类型元素的数组
type Element interface {
}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

//复制数据切片至空接口切片

/*var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for i, d := range dataSlice {
interfaceSlice[i] = d
}*/

//接口到接口
//只要被转换的接口中包含了全部转换接口的方法，则可以转换成功，否则会出现运行时的异常
//任何接口都可以转换给空接口
/*type myPrintInterface interface {
	print()
}

func f3(x myInterface) {
	x.(myPrintInterface).print() // type assertion to myPrintInterface
}*/
