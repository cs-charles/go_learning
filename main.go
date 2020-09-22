package main

import (
	"github.com/dscxieyong/go_learning/const_package"
	"github.com/go_learning/interface_package"
)

func main() {
	//常量部分
	const_package.PrintUppercase()
	//const_package.PrintLowercase()
	//const_package.PrintNumber()

	//变量
	//variable_package.F1()
	//variable_package.F2()

	//数据类型与操作符
	//operation_sign_package.Operation()

	//字符串
	//string_package.PrintStr()
	//string_package.StringsFun()
	//string_package.StringConv()

	//日期时间
	//date_time_package.Time()

	//指针
	//point_package.PrintPoint()
	//point_package.ChangeValue()

	//结构控制
	/*construct_package.UseIf()
	construct_package.UseSwitch(5)
	construct_package.UseFor()*/

	//函数
	/*appendJPG := function_package.AppendEndFix(".jpg")
	fmt.Printf("拼接后缀jpg: %s \n",appendJPG("照片"))
	appendPDF := function_package.AppendEndFix(".pdf")
	fmt.Printf("拼接后缀pdf: %s \n",appendPDF("文档"))
	function_package.CalculationUseTime()*/
	/*function_package.GoRecoverPanic()
	function_package.MainRecoverPanic()*/

	//数组与切片
	/*arr_with_slice_package.InterArr()
	arr_with_slice_package.CopyArr1()
	arr_with_slice_package.CopyArr2()
	arr_with_slice_package.CopyArr3()
	arr_with_slice_package.SliceCopy()
	arr_with_slice_package.Reslice()
	arr_with_slice_package.Copy()
	arr_with_slice_package.Append()
	arr_with_slice_package.GrowthCap(make([]int,10,10),10)
	arr_with_slice_package.Converter()
	arr_with_slice_package.ChangeString()*/

	//内置包
	/*pack_package.Match("John: 2578.34 William: 4567.23 Steve: 5632.18")
	pack_package.Replace("John: 2578.34 William: 4567.23 Steve: 5632.18")
	pack_package.ReplaceByFunction("John: 2578.34 William: 4567.23 Steve: 5632.18")
	pack_package.BigIntCal()
	pack_package.BigRatCal()
	pack_package.TestFloatCal()*/

	//结构体和方法
	//struct_method_package.InitStructValue()
	//struct_method_package.Converter()
	//struct_method_package.Compare()
	//struct_method_package.PrintlnAnonymous()
	//struct_method_package.Cal()
	//struct_method_package.Finalizer()

	//interface_package.TestInterface()
	//interface_package.InterfaceAsset()
	//interface_package.TypeSwitch()
	//interface_package.SortStrings([]string{"1","2"})
	//interface_package.TestTypeOf()
	//interface_package.TestValueOf()
	//interface_package.TestSetValue()
	//interface_package.TestSetStructValue()

	//文件读写
	//read_write_package.ReadFromStdIn()
	//read_write_package.ReadFile1()
	//read_write_package.ReadFile2()
	//read_write_package.CopyFile()
	//read_write_package.ReadCol()
	//read_write_package.FilePath()
	//read_write_package.WriteFile1()
	//read_write_package.WriteFile2()
	//read_write_package.PrintByFlag()
	//read_write_package.TestCat()
	//read_write_package.Marshal()
	//read_write_package.UnMarshal()

	//goroutine
	//goroutine_package.TestSemaphore()
	//goroutine_package.TestSelect()
	//goroutine_package.Ticker()
	//goroutine_package.TestMultiplexing()

	//net
	//web.TestHead()
	//web.StartSimpleWebServer()

	//context
	//context.TestContext()
	//str := fmt.Sprintf("%v--%v","1231",1213)
	//fmt.Println(str)
	interface_package.TestInterfaceWithMethod()
}
