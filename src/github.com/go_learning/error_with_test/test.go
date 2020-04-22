package error_with_test

// Go 的测试工具 gotest
//测试程序必须属于被测试的包，并且文件名满足这种形式 *_test.go，所以测试代码和包中的业务代码是分开的。
//测试文件中必须导入 "testing" 包，并写一些名字以 TestZzz 打头的全局函数，这里的 Zzz 是被测试函数的字母描述，如 TestFmtInterface，TestPayEmployees 等。

//func TestAbcde(t *testing.T)

//T 是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如 t.Log，t.Error，t.ErrorF 等。在函数的结尾把输出跟想要的结果对比，如果不等就打印一个错误。成功的测试则直接返回。

//用下面这些函数来通知测试失败：
//
//1）func (t *T) Fail()
//
//    标记测试函数为失败，然后继续执行（剩下的测试）。
//2）func (t *T) FailNow()
//
//    标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
//3）func (t *T) Log(args ...interface{})
//
//    args 被用默认的格式格式化并打印到错误日志中。
//4）func (t *T) Fatal(args ...interface{})
//
//    结合 先执行 3），然后执行 2）的效果。
//

//gotest 可以接收一个或多个函数程序作为参数，并指定一些选项。
//
//结合 --chatty 或 -v 选项，每个执行的测试函数以及测试状态会被打印。
//例如：go test fmt_test.go --chatty

//基准测试

//；测试代码中必须包含以 BenchmarkZzz 打头的函数并接收一个 *testing.B 类型的参数，比如：
//
//func BenchmarkReverse(b *testing.B) {
//    ...
//}

//命令 go test –test.bench=.* 会运行所有的基准测试函数；代码中的函数会被调用 N 次（N 是非常大的数，如 N = 1000000），并展示 N 的值和函数执行的平均时间，单位为 ns（纳秒，ns/op）。

//一个测试函数至少包含三个测试用例
//正常的用例
//反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
//边界检查用例（如果参数的取值范围是 0 到 1000，检查 0 和 1000 的情况）
