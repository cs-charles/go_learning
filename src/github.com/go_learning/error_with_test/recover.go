package error_with_test

import "fmt"

//recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。

//总结：panic 会导致栈被展开直到 defer 修饰的 recover () 被调用或者程序中止。

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}
