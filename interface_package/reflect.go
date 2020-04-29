package interface_package

import (
	"fmt"
	"reflect"
)

//反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如它的大小、方法和 动态
//的调用这些方法。这对于没有源代码的包尤其有用

//变量的最基本信息就是类型和值：
//反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。

//两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，
//返回被检查对象的类型和值。
//例如，x 被定义为：var x float64 = 3.4，那么 reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 <float64 Value>
/**
func TypeOf(i interface{}) Type
func ValueOf(i interface{}) Value
*/

//Kind()方法可以返回底层的值类型

type BizInt int

var myInt BizInt = 5

func TestTypeOf() {
	t := reflect.TypeOf(myInt)
	fmt.Println("myInt kind is", t.Kind())
	fmt.Println("myInt type is", t)

}

func TestValueOf() {
	v := reflect.ValueOf(myInt)
	fmt.Println("myInt value is", v)
	fmt.Println("myInt value kind is", v.Kind())
	fmt.Println("myInt value type is", v.Type())
	fmt.Println("myInt value is:", v.Int())
	fmt.Println("myInt value is", v.Interface())
	//接口类型断言
	fmt.Println("myInt value is", v.Interface().(BizInt))
}

func TestSetValue() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	//如果v的Kind不是Interface或Ptr会panic；
	//如果v持有的值为nil，会返回Value零值。
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}

//struct set value
type NotknownType struct {
	S1, S2, S3 string
}

type ChangeValuer interface {
	changeValue(s string)
}

func (n NotknownType) String() string {
	return n.S1 + " - " + n.S2 + " - " + n.S3
}

func TestSetStructValue() {
	//只有NotknownType指针才可以更改NotknownType的值
	var secret interface{} = &NotknownType{"Ada", "Go", "Oberon"}
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)
	//只有执行了Elem()方法才可以设置value
	value = value.Elem()
	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//结构体中只有大写的被导出的属性才可以被设置值,并且Value类型必须满足是持有类型的指针reflect.ValueOf(&secret)
		value.Field(i).SetString("C#")
	}
	//构造一个类型的值对象
	//reflect.New(type)
	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
