package goroutine_package

import (
	"fmt"
	"time"
)

//在协程周期性的执行一些事情（打印状态日志，输出，计算等等）的时候非常有用

/**
type Ticker struct {
    C <-chan Time // the channel on which the ticks are delivered.
    // contains filtered or unexported fields
    ...
}
*/
/**
时间间隔的单位是 ns（纳秒，int64），在工厂函数 time.NewTicker 中以 Duration 类型的参数传入：func Newticker(dur) *Ticker。
在协程周期性的执行一些事情（打印状态日志，输出，计算等等）的时候非常有用。
调用 Stop() 使计时器停止，在 defer 语句中使用
*/

//time.Tick() 函数声明为 Tick(d Duration) <-chan Time，
//当你想返回一个通道而不必关闭它的时候这个函数非常有用：

func Ticker() {
	go func() {
		ticker := time.NewTicker(1e9)
		defer ticker.Stop()
		for {
			v := <-ticker.C
			fmt.Printf("ticker return value is %v\n", v)
		}
	}()
	go func() {
		//time.Tick与time.NewTicker的区别是time.Tick无需关闭
		ch := time.Tick(1e9)
		time := <-ch
		fmt.Printf("ticker return value is %v\n", time.String())
	}()
	time.Sleep(1e10)
}

//定时器（Timer）结构体看上去和计时器（Ticker）结构体的确很像（构造为 NewTimer(d Duration)），
//但是它只发送一次时间，在 Dration d 之后。

//time.After(d)在 Duration d 之后，当前时间被发到返回的通道；所以它和 NewTimer(d).C 是等价的

//timer配合select可以实现超时模式
func timeOutRead(second int64) (ch chan bool) {
	ch = make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(second))
		ch <- true
	}()
	return
}

func ReadFromChannel(ch chan bool, timeout chan bool) {
	for {
		select {
		case <-ch:
			// a read from ch has occured
		case <-timeout:
			// the read from ch has timed out
			break
		}
	}
}

/**
超时设置2
*/

/**
ch := make(chan error, 1)
go func() { ch <- client.Call("Service.Method", args, &reply) } ()
select {
case resp := <-ch
    // use resp and reply
case <-time.After(timeoutNs):
    // call timed out
    break
}
*/

/**
//从数据库切片中读取数据
func Query(conns []conn, query string) Result {
    ch := make(chan Result, 1)
    for _, conn := range conns {
        go func(c Conn) {
            select {
            case ch <- c.DoQuery(query):
            default:
            }
        }(conn)
    }
    return <- ch
}
*/

/**
在应用中缓存数据：
应用程序中用到了来自数据库（或者常见的数据存储）的数据时，经常会把数据缓存到内存中，
因为从数据库中获取数据的操作代价很高；如果数据库中的值不发生变化就没有问题。
但是如果值有变化，我们需要一个机制来周期性的从数据库重新读取这些值：缓存的值就不可用（过期）了，
而且我们也不希望用户看到陈旧的数据。
*/

//执行的协程可以总是可以使用 runtime.Goexit() 来停止。
