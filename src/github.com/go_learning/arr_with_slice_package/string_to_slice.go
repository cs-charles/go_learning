package arr_with_slice_package

import (
	"fmt"
	"unicode/utf8"
)

//统计字符串中字符的数目utf8.RuneCountInString
//字符串转换成slice,字符串本质上也是一个字节数组
//通过 []byte() 或者 []int32
func Converter() {
	str1 := "ABCDEFG"
	str2 := "我爱中国🇨🇳ABCD"
	fmt.Printf("字符串数目,str1 count is %d,str2 count is %d \n", utf8.RuneCountInString(str1), utf8.RuneCountInString(str2))
	slice1 := []byte(str1)
	slice2 := []int32(str2)
	fmt.Printf("%c \n", slice1)
	fmt.Printf("%c \n", slice2)
}

func ChangeString() {
	str1 := "01234"
	slice1 := []byte(str1)
	fmt.Printf("old string is %s ,old slice is %c\n", str1, slice1)
	//更改slice值是不会影响到str1,这是golang为了保护字符串安全
	slice1[0] = 'A'
	//更改字符串str1
	str1 = string(slice1)
	fmt.Printf("old string is %s ,old slice is %c\n", str1, slice1)
}
