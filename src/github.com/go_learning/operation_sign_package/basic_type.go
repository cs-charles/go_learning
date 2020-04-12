package operation_sign_package

//类型转换和赋值操作
var i1 = int64(1002)

//go不同类型之间必须通过显示转换才可以运算
var i2 = int32(i1)

//进制表示 0表示8进制，0x表示16进制,e表示10的连乘
var i3 = 011

var i4 = 0x001

var i5 = 2.3e3

//int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
//uintptr 的长度被设定为足够存放一个指针即可。

func PrintNumber() {
	println(i1, i2, i3, i4, i5)
}

//格式化输出
// d% 10进制表示
// x% 小写16进制数字
// X% 大写16进制数字
// g% 后面不带小数的浮点数
// f% 后面带小数的浮点数
// e% 科学计数法表示
// %n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，
//例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00。

//运算符
// & | ^ 且 或 异或
// &^ 将运算符左边数据相异的位保留，相同位清零。如果右边是0则左边不变，如果右边是1，则左边清零
// ^m 求m的补码 如下面的列子 m=-4 m的补码就是0000 0011
func Operation() {
	a := 100
	b := 4
	var c int8 = -4
	println(a &^ 1)
	println(b &^ 3)
	println(^c)
	println(ch2 == ch3)
	println(ch4)
}

//类型别名 类型别名得到的新类型并非和原类型完全相同，新类型不会拥有原类型所附带的方法
type TZ int
type REPO string

var i6 TZ
var str REPO

// byte是uint8的别名 通俗的讲byte代表 Java中的char表示的ASCII码值
// rune是int32的别名 rune代表Java中所有表示的char类型
var ch0 byte = 'A'
var ch1 byte = 65
var ch2 byte = '\x41'
var ch3 byte = 0x41
var ch4 rune = '我'

//var ch5 byte = '我'  //会报错，因为byte只能表示ASCII码值

//（\x 总是紧跟着长度为 2 的 16 进制数）
//另外一种可能的写法是 \ 后面紧跟着长度为 3 的八进制数，例如：\377。
//在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U。
//因为 Unicode 至少占用2个字节，所以我们使用int16 或者int 类型来表示。
//如果需要使用到 4 字节，则会加上 \U 前缀；前缀 \u 则总是紧跟着长度为 4 的 16 进制数，
//前缀 \U 紧跟着长度为 8 的 16 进制数。
