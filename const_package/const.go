package const_package

const A = "A"
const (
	C = "C"
	B = "B"
)
const one, two, three = 1, 2, 3
const (
	a = iota
	b
	c
)

const (
	e = iota + 1
	f = iota + iota
	g = iota * 2
)

func PrintUppercase() {
	//允许在块级作用域内声明重复的常量A
	const A = 1
	//常用不允许采用初始化赋值
	//const_package B := 2
	println(A, B, C)
}

func PrintLowercase() {
	println(a)
	println(b)
	println(c)
	println(e)
	println(f)
	println(g)
}

func PrintNumber() {
	println(one)
	println(two)
	println(three)
}
