package goroutine_package

import (
	"strconv"
)

//为了保护一个对象的并发修改，我们可以使用一个后台的协程来顺序执行一个匿名函数，
//而不是通过同步 互斥锁（Mutex） 进行锁定。

type People struct {
	name   string
	salary float64
	ch     chan func()
}

func NewPeople() *People {
	p := &People{"陈三", 53.243, make(chan func())}
	go p.execute()
	return p
}

func (p *People) execute() {
	for f := range p.ch {
		f()
	}
}

func (p *People) setSalary(salary float64) {
	p.ch <- func() {
		p.salary = salary
	}
}

func (p *People) Salary() float64 {

	fChan := make(chan float64)

	p.ch <- func() { fChan <- p.salary }

	return <-fChan

}

func (p *People) String() string {

	return "Person - name is: " + p.name + " - salary is: " +

		strconv.FormatFloat(p.Salary(), 'f', 2, 64)

}
