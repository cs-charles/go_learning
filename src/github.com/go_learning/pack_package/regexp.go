package pack_package

import (
	"fmt"
	"regexp"
	"strconv"
)

var pattern = "[0-9]+.[0-9]+" //正则

//Compile在编译正则时，如果出错会返回error,如果采用MustCompile一旦出现程序会出现panic
//在程序中出现panic会让程序立马退出(除非recover)
//recover：原意重新获得、恢复。上面说到panic函数并不会立刻返回，而是先去执行defer，然后再返回。这时候如果在defer中阻止panic继续向上传递，那就异常的处理机制完善了。
//go语言提供了recover内置函数，前面提到，在defer中可以使用recover来捕获panic，禁止panic向上传递，从而不会导致整个程序的退出。

var reg, _ = regexp.Compile(pattern)

func f(s string) string {
	v, _ := strconv.ParseFloat(s, 32)
	return strconv.FormatFloat(v*2, 'f', 2, 32)
}

//正则匹配
func Match(str string) {
	bool, _ := regexp.Match(pattern, []byte(str))
	if bool {
		fmt.Printf("match \n")
	} else {
		fmt.Printf("not match \n")
	}
}

//替换
func Replace(str string) {
	byte := reg.ReplaceAll([]byte(str), []byte("###.#"))
	fmt.Printf("after replace value is %s\n", byte)
}

//方法替换,根据方法返回内容替换
func ReplaceByFunction(str string) {
	byte := reg.ReplaceAllStringFunc(str, f)
	fmt.Printf("after replace value is %s\n", byte)
}
