package error_with_test

import (
	"errors"
	"fmt"
)

//永远不要忽略错误，否则可能会导致程序崩溃！！
//defer-panic-and-recover机制

//errors 包中有一个 errorString 结构体实现了 error 接口。当程序处于错误状态时可以用 os.Exit(1) 来中止运行。
/*type error interface {
	Error() string
}
*/

func TestError() {
	err := errors.New("这是一个自定义错误")
	fmt.Println(err.Error())
}

//在大部分情况下自定义错误结构类型很有意义的，可以包含除了（低层级的）错误信息以外的其它有用信息，
//例如，正在进行的操作（打开文件等），全路径或名字。看下面例子中 os.Open 操作触发的 PathError 错误
/*type PathError struct {
   Op string    // "open", "unlink", etc.
   Path string  // The associated file.
   Err error  // Returned by the system call.
}

func (e *PathError) String() string {
   return e.Op + " " + e.Path + ": "+ e.Err.Error()
}*/

//错误类型判断
/**
if e, ok := err.(*os.PathError); ok {
    // remedy situation
}
*/

//fmt 创建错误对象
//fmt.Errorf("math: square root of negative number %g", f)
