package goroutine_package

import (
	"fmt"
	"sync"
	"time"
)

//old model
type Task struct {
	//some things
}

type TaskPool struct {
	lock  sync.Mutex
	tasks []Task
}

func SyncWorker(pool *TaskPool) {
	for {
		//加锁
		pool.lock.Lock()
		//获取任务
		task := pool.tasks[0]
		//更新任务池
		pool.tasks = pool.tasks[1:]
		fmt.Println(task)
		//解锁
		pool.lock.Unlock()
	}

}

//sync.Mutex(参见 9.3 是互斥锁：它用来在代码中保护临界区资源：同一时间只有一个 go 协程（goroutine）可以进入该临界区。
//如果出现了同一时间多个 go 协程都进入了该临界区，则会产生竞争

//new model
func SendTask(in chan *Task) {
	for {
		task := new(Task)
		in <- task
		time.Sleep(1e10)
	}
}

func ChannelWork(in, out chan *Task) {
	for {
		task := <-in
		//处理task
		fmt.Println(task)
		out <- task
	}
}

func ConsumerWork(in chan *Task) {
	//....
}

func TestNewWork(n int) {
	pending, done := make(chan *Task), make(chan *Task)
	go SendTask(pending)
	for i := 0; i < n; i++ {
		go ChannelWork(pending, done)
	}
	ConsumerWork(pending)
}

/**
两者对比
使用锁的情景：
访问共享数据结构中的缓存信息
保存应用程序上下文和状态信息数据

使用通道的情景：
与异步操作的结果进行交互
分发任务
传递数据所有权
*/
