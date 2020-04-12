package variable_package

var a string = "a"

func F1() {
	{
		//局部变量a的初始化不会影响全局变量a
		a := "O"
		println(a)
	}
	//输出全局变量a
	a = "b"
	println(a)
}

func F2() {
	println(a)
}
