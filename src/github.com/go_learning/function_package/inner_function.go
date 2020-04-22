package function_package

import (
	"fmt"
	"time"
)

//内置函数
// close 用于管道通信
// len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）
// cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
// new new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，
// make 用于内置引用类型（切片、map 和管道）分配内存
// copy 用于复制切片(slice)
// append 用于连接切片
// panic、recover	两者均用于错误处理机制
// print、println	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
// complex、real imag	用于创建和操作复数（详见第 4.5.2.2 节）

//golang的一个协程异常，可以通过defer捕获
func GoRecoverPanic() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("has recover panic \n", r)
			}
		}()
		time.Sleep(time.Second * 2)
		panic(1)
	}()
	fmt.Printf("main thread \n")
}

//协程中如果没有捕获，回导致整个进程退出。
func MainRecoverPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has recover panic \n", r)
		}
	}()
	go func() {
		fmt.Printf("coroutine has panic\n")
		time.Sleep(time.Second * 2)
		panic(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(time.Second * 4)
	}

	fmt.Printf("main thread \n")
}
