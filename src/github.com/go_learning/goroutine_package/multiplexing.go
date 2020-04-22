package goroutine_package

import (
	"fmt"
	"time"
)

//io多路复用
type Response struct {
	res int
}

type Request struct {
	arg1     int
	arg2     int
	response chan *Response
}

type Operation func(int, int) int

func server(op Operation, server chan *Request) {
	for {
		req := <-server
		// 为请求开一个 goroutine
		go func() {
			res := op(req.arg1, req.arg2)
			response := &Response{res: res}
			req.response <- response
		}()
	}
}

func startServer(op Operation) chan *Request {
	req := make(chan *Request)
	go server(op, req)
	return req
}

func TestMultiplexing() {
	start := time.Now()
	adder := startServer(func(a, b int) int { return a + b })
	const N = 10000

	var reqs [N]Request

	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.arg1 = i
		req.arg2 = i + N
		req.response = make(chan *Response)
		//发送请求
		adder <- req
	}

	//校验
	for i := 0; i < N; i++ {
		//堵塞，直到收到通道响应
		if response := <-reqs[i].response; response.res != 2*i+N {
			fmt.Println("request ", i, " has fault")
		} else {
			fmt.Println("request ", i, " has success")
		}
	}

	end := time.Now()

	fmt.Printf("use time is %s", end.Sub(start))
}
