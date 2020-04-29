package goroutine_package

import (
	"flag"
	"fmt"
)

//链式操作

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

//内容从通道in输入到out
func f(out, in chan int) {
	out <- 1 + <-in
}

func Chain() {
	flag.Parse()
	leftMost := make(chan int)
	var left, right chan int = nil, leftMost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 1
	out := <-leftMost
	fmt.Println(out)
}
