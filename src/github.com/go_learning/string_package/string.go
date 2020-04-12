package string_package

import (
	"fmt"
	"strings"
)

//字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）
//字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容；更深入地讲，字符串是字节的定长数组。

/*解释字符串：
该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：

\n：换行符
\r：回车符
\t：tab 键
\u 或 \U：Unicode 字符
\\：反斜杠自身

非解释字符串：
该类字符串使用反引号括起来，支持换行*/

var str1 = "hello world1"
var str2 = `hello` +
	`world2`

func PrintStr() {
	println(str1)
	println(str2)
	//第一个字符
	s0 := str1[0]
	sn := str1[len(str1)-1]
	fmt.Printf("%c \n", s0)
	fmt.Printf("%c", sn)
}

func StringsFun() {
	//判断前缀
	fmt.Printf("prefix is hello %t \n", strings.HasPrefix(str1, "hello"))
	//判断后缀
	fmt.Printf("end prefix is world1 %t \n", strings.HasSuffix(str2, "world1"))
	//判断包含
	fmt.Printf("contains hello? %t \n", strings.Contains(str1, "hello"))

}
