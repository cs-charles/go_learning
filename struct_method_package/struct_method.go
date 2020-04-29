package struct_method_package

import (
	"fmt"
	"strconv"
)

//Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
//一个类型加上它的方法等价于面向对象中的一个类。
//一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。

//格式定义
/*func (recv recvType )(param_list) (result_list)  {

}*/
//因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。
//但是如果基于接收者类型，是有重载的：
//具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的：
/**
func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix
*/

//如果给指定类型指定别名，则可以实现在不同包下定义方法，但是方法必须通过别名访问

//go在方法调用时存在自动解引用，和自动取引用。如果方法的recv是指针类型的话，则可以对原值产生影响。
//但是在如果存在接口的方法类型匹配时会有不同的情况发生。
//值类型的方法总会生成一个对应的指针类型方法，反之则不可以。

type employee struct {
	salary float64
}

func (e *employee) giveRaise(rate float64) {
	e.salary = (1 + rate) * e.salary
}

type myArr []int

func (e myArr) Inter() {
	for _, v := range e {
		fmt.Printf("position i value is %d\n", v)
	}
}

func Cal() {
	e := employee{salary: 15000.19}
	e.giveRaise(0.1)
	fmt.Printf("new salary is %g\n", e.salary)
	arr := myArr{1, 2}
	arr.Inter()
}

//方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

//对于未导出的属性，可以通过getter,setter方法获取或改变

//内嵌类型的方法和继承
//内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法 “晋升” 成为了外层类型的方法。
//当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，
//可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。

//String() 自定义结构体的格式化方式，fmt.Print会调用此方法。

func (e *employee) String() string {
	//f代表格式化占位符，精度是2，float64
	return "(" + strconv.FormatFloat(e.salary, 'f', 2, 64)
}
