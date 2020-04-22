package string_package

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringConv() {
	//int to string
	s := strconv.Itoa(64)
	fmt.Printf("64 string type is %v\n", s)
	//string to int
	i, _ := strconv.Atoi("12")
	fmt.Printf("i value is %v\n", i)

	//parse类型
	//bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8。
	//base参数表示以什么进制的方式去解析给定的字符串，有效值为0、2-36。当base=0的时候，表示根据string的前缀来判断以什么进制去解析：0x开头的以16进制的方式去解析，0开头的以8进制方式去解析，其它的以10进制方式解析。
	bool, _ := strconv.ParseBool("t")
	fmt.Printf("bool value is %v\n", bool)
	f, _ := strconv.ParseFloat("23.000242", 64)
	fmt.Printf("f value is %v\n", f)
	i0, _ := strconv.ParseInt("16", 16, 8)
	fmt.Printf("i0 value is %v\n", i0)
	i1, _ := strconv.ParseUint("1000000", 10, 8)
	fmt.Printf("i1 value is %v\n", reflect.TypeOf(i1))

	//format类型
	s0 := strconv.FormatBool(true)
	fmt.Printf("s0 value is %v\n", s0)
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	s1 := strconv.FormatFloat(124.020, 'f', 2, 32)
	fmt.Printf("s1 value is %v\n", s1)
	s2 := strconv.FormatInt(1111, 2)
	fmt.Printf("s2 value is %v\n", s2)
	s3 := strconv.FormatUint(222, 16)
	fmt.Printf("s3 value is %v\n", s3)

	//append类
	baseByte := []byte("base byte")
	destByte := strconv.AppendBool(baseByte, true)
	fmt.Printf("destByte value is %v\n", string(destByte))
	destByte = strconv.AppendFloat(baseByte, 4324.23, 'f', 3, 32)
	fmt.Printf("destByte value is %v\n", string(destByte))
	destByte = strconv.AppendInt(baseByte, 23424, 2)
	fmt.Printf("destByte value is %v\n", string(destByte))
	destByte = strconv.AppendUint(baseByte, 245234, 2)
	fmt.Printf("destByte value is %v\n", string(destByte))
}
