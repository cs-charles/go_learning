package function_package

import "fmt"

//go内部有三种函数
//普通的带有名字的函数
//匿名函数或者 lambda 函数
//方法Methods
//golang不允许函数重载，不支持这项特性的主要原因是函数重载需要进行多余的类型匹配影响性能

//非命名返回值
func F1(input int) (int, int) {
	return input * 5, input * 10
}

//命名返回值
func F2(input int) (x1 int, x2 int) {
	x1, x2 = input*5, input*10
	return
}

//传递变长参数
func max(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	maxVal := s[0]
	for _, val := range s {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

//多个不确定类型的参数采用interface
func typeCheck(values ...interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Printf("int type")
		case float64:
			fmt.Printf("float type")
		case string:
			fmt.Printf("string type")
		case bool:
			fmt.Printf("boolean type")
		}
	}
}

//关键字 defer 类型Java中的finally，一般用于释放资源，如果有多个defer先申明的后执行
//常见用途
//1.关闭文件流
//2.解锁一个加锁的资源
//3.打印最终报告
//4.关闭数据库链接
//5.可以更改函数命名返回值
func deferFun() {
	println("start to do")
	defer println("释放资源")
	println("end to do")
}
