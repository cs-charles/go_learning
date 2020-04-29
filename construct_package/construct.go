package construct_package

import "fmt"

func UseIf() {
	if count, error := fmt.Printf("嗨咯"); error != nil {
		fmt.Printf("打印失败，打印字符数目:%d", count)
	} else {
		fmt.Printf("打印成功，打印字符数目:%d", count)
	}
}

func UseSwitch(i int) {
	switch i {
	case 10:
		fmt.Printf("i = 10 \n")
		fallthrough
	case 8:
		fmt.Printf("i = 8 \n")
	default:
		fmt.Printf("default \n")
	}
}

func UseFor() {
	i := 0
	for i < 10 {
		fmt.Printf("i = %d\n", i)
		i++
	}

	for i, j := 0, 0; i < 5; {
		fmt.Printf("i = %d , j = %d \n", i, j)
		i++
		j++
	}

	//for-range
	var str = "中国话"
	for inx, val := range str {
		fmt.Printf("index is %d ,value is %c\n", inx, val)
	}

	//label
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				//当执行到该语句的时候，就会跳转到 LABEL1 标签的位置执行continue，所以导致j==4和j==5都不会被执行，实际上是continue掉了外层的for
				//退出内层循环，如果为break LABEL1则退出内外两层循环
				//goto LABEL1，跳转至LABEL1继续执行此段代码
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
