package interface_package

//显式地指明类型实现了某个接口
//通过指定某个特殊的方法实现例如
type Swing interface {
	Swing()
	ImplementsSwing()
}

//方法重载，在go语言中不支持方法的重载，不过可以通过空接口来实现

//当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。
