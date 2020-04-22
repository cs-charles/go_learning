package goroutine_package

import "bytes"

//漏桶算法

//  ！！！在一个select语句中，Go语言会按照顺序从头到尾评估每一个发送和接收
//的语句。如果其中任意一语句可以继续执行（即没有被阻塞），那么就那些可以执行的
//语句中任意选择一条来使用。

//存储buffer
var freeList = make(chan *bytes.Buffer, 100)

//传递buffer
var serverChan = make(chan *bytes.Buffer)

func client() {

	for {

		var b *bytes.Buffer

		// 如果 freeList 通道中有 buffer，直接获取；如果没有，就创建一个新的

		select {

		case b = <-freeList:

		// 获取到一个 ，没有做其他事情

		default:

			// 没有空闲的，所以分配一个新的

			b = new(bytes.Buffer)

		}

		/*loadInto(b) // 从网络去获取下一条信息

		serverChan <- b // 发送给服务器端*/

	}

}

func servers() {

	for {

		b := <-serverChan // 等待工作。（等待客户端发送一个 buffer 过来）

		//process(b)

		// 如果就空间，就重用 buffer

		select {

		case freeList <- b:

			// 如果 freeList 有空闲的插槽，就重用 buffer；没有做其他事情

		default:

			// freeList 已满，只是继续： 会将 buffer 掉落（丢弃）

		}

	}

}
