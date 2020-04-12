package main

import (
	"github.com/go_learning/const_package"
	"github.com/go_learning/operation_sign_package"
	"github.com/go_learning/string_package"
	"github.com/go_learning/variable_package"
)

func main() {
	//常量部分
	const_package.PrintUppercase()
	const_package.PrintLowercase()
	const_package.PrintNumber()

	//变量
	variable_package.F1()
	variable_package.F2()

	//数据类型与操作符
	operation_sign_package.Operation()

	//字符串
	string_package.PrintStr()
	string_package.StringsFun()

	//
}
