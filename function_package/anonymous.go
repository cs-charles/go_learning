package function_package

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//函数递归栈溢出的问题，可以通过懒惰求值的方式解决

func init() {
	log.SetFlags(log.Llongfile)
}

//函数作为入参
func FunAsArg(f func(i int, j int) bool) int {
	bool := f(1, 1)
	if bool {
		return 1 * 1
	}
	return 0
}

//匿名函数
var f = func(x, y int) int {
	return x + y
}

func FuncNoName(x, y int) int {
	if x > 0 && y > 0 {
		//匿名函数用法1
		return f(x, y)
	}
	//匿名函数用法2
	return func(x, y int) int { return x * y }(x, y)
}

//一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，
//因此这些变量可以在闭包中被操作，直到被销毁
//闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
func Adder() func(int) int {
	//下面这个x属性在每一次调用时都会把值保存下来
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

//闭包也可以用于修改外部变量的值
var G int

func ModifyG() {
	//编译器会把 go 后面跟着的参数与函数都打包成了对象，等待系统调度。
	//具体见：https://studygolang.com/articles/16469?fr=sidebar
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		G = s
	}(1000)
}

//部分求值
func AppendEndFix(endFix string) func(str string) string {
	where(endFix)
	return func(name string) string {
		if strings.HasSuffix(name, endFix) {
			return name
		}
		where(name)
		return name + endFix
	}
}

//闭包调试
//方式一，采用runtime.Caller()
/*var where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}*/

//方式二 log.SetFlags(log.Llongfile)在init中设置
var where = log.Print

//计算函数耗时
func CalculationUseTime() {
	start := time.Now()
	FuncNoName(100, 100)
	end := time.Now()
	fmt.Printf("longCalculation took this amount of time: %s ", end.Sub(start))
}
