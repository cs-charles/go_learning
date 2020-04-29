package goroutine_package

import (
	"fmt"
	"time"
)

//协程可以使用共享变量来通信，但是很不提倡这样做，因为这种方式给所有的共享内存的多线程都带来了困难。
//而 Go 有一个特殊的类型，通道（channel），像是通道（管道），可以通过它们发送类型化的数据在协程之间通信，
//可以避开所有内存共享导致的坑；通道的通信方式保证了同步性。数据通过通道：同一时间只有一个协程可以访问数据：所以不会出现数据竞争，设计如此。数据的归属（可以读写数据的能力）被传递。

//var identifier chan datatype
//未初始化的通道的值是 nil。
//
//所以通道只能传输一种类型的数据，比如 chan int 或者 chan string，
//所有的类型都可以用于通道，空接口 interface{} 也可以。甚至可以（有时非常有用）创建通道的通道。

var ch1 chan string

//初始化
//ch1 = make(chan string)

//函数通道
//funcChan := chan func()

/**
流向通道（发送）

ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

从通道流出（接收），三种方式：

int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）
；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。

<- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：

if <- ch != 1000{
    ...
}
*/

func TestChannel() {
	var ch = make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "hello world"
	ch <- "你好，世界"
	ch <- "哈哈哈"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}

// 通道阻塞
//默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。
//可以想象一个无缓冲的通道在没有空间来保存数据的时候：
//必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。
//所以通道的发送 / 接收操作在对方准备好之前是阻塞的：

//）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果 ch 中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。
//
//2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

//通过一个（或多个）通道交换数据进行协程同步
func f1(in chan int) {
	fmt.Println(<-in)
}

func DeadLock() {
	out := make(chan int)
	//此处会导致一直堵塞，所以协程无法启动
	out <- 2
	go f1(out)
}

// 同步通道 - 使用带缓冲的通道
//发送在缓冲满之前都不会堵塞，接收在通道容量为0之前也都不会堵塞
var buf = 100
var ch2 = make(chan string, buf)

//用带缓冲通道实现一个信号量

//信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的 sync 的 Go 包，使用带缓冲的通道可以轻松实现：
//
//带缓冲通道的容量和要同步的资源容量相同
//通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
//容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）

//通道工厂模式，在生成通道的内部有个协程输入数据
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//给通道使用 for 循环
/**
for v := range ch {
    fmt.Printf("The value is %v\n", v)
}
*/

//迭代器模式
/**
func (c *container) Iter () <- chan items {
    ch := make(chan item)
    go func () {
        for i:= 0; i < c.Len(); i++{    // or use a for-range loop
            ch <- c.items[i]
        }
    } ()
    return ch
}

for x := range container.Iter() { ... }
*/

//通道的方向
var send_only chan<- int // channel can only receive data
var recv_only <-chan int // channel can only send data
//只接收的通道无法关闭
/**
var c = make(chan int) // bidirectional
go source(c)
go sink(c)

func source(ch chan<- int){
    for { ch <- 1 }
}

func sink(ch <-chan int) {
    for { <-ch }
}
*/

//管道和选择器模式
/**
sendChan := make(chan int)
reciveChan := make(chan string)
go processChannel(sendChan, receiveChan)

func processChannel(in <-chan int, out chan<- string) {
    for inValue := range in {
        result := ... /// processing inValue
    out <- result
    }
}
*/

//通道关闭
//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。
//通过close(ch)

//用来检测通道没有被关闭，且没有堵塞
/**
if v, ok := <-ch; ok {
  process(v)
}
*/

//for -range 会自动检测通道是否被关闭
/**
for input := range ch {
    process(input)
}
*/
