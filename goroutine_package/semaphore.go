package goroutine_package

import (
	"fmt"
	"time"
)

type Empty interface{}
type Semaphore chan Empty

func (s Semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func TestSemaphore() {
	n := 10
	s := make(Semaphore, n)
	go func() {
		s.P(10)
		fmt.Println("go1 sleep start")
		time.Sleep(2e9)
		s.V(10)
		fmt.Println("go1 sleep end")
	}()

	go func() {
		s.P(1)
		fmt.Println("go2 sleep start")
		time.Sleep(2e9)
		s.V(1)
		fmt.Println("go2 sleep end")
	}()
	time.Sleep(10e9)
}
