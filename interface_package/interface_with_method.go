package interface_package

//使用方法集与接口
//如果实现接口的方法定义在指针
//在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
//
//指针方法可以通过指针调用
//值方法可以通过值调用
//接收者是值的方法可以通过指针调用，因为指针会首先被解引用
//接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

/***
Go 语言规范定义了接口方法集的调用规则：
类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
类型 T 的可调用方法集包含接受者为 T 的所有方法
类型 T 的可调用方法集不包含接受者为 *T 的方法***/

/**
理解上述规则假设存在两个方法
func (a *int)add(b int){
	*a += b
}
如果自动生成方法则会是
func (a int)add(b int){
	//显然只是会改变形参a,对于实际调用此方法的a对象并不会有影响
	&a.add(b)
}

假设存在方法
func (a int)add(b int){
	a += b
	...
}
则会自动生成
func (a *int)add(b int){
	//此处同样只是值传递，不会影响到外部的对象，所以合理
	*a.add(b)
}
*/

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func TestInterfaceWithMethod() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	//CountInto(lst, 1, 10) //不被允许
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
