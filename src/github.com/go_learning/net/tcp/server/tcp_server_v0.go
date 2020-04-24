package main

import (
	"fmt"
	"net"
)

/*type Clients struct {
	clients []*net.Conn
	cursor int
}

func NewClients() *Clients {
	clients := new(Clients)
	return clients
}*/

//net 包中网络通信的功能。它包含了用于 TCP/IP 以及 UDP 协议、域名解析等方法。
func StartServer() {
	fmt.Println("server starting ....")
	//创建listener
	listen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	//监听请求
	for {
		//堵塞
		con, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		go doServerStuff(con)
	}
	//比较两个对象是否相等
	//reflect.DeepEqual()
}

func doServerStuff(con net.Conn) {
	buf := make([]byte, 1024)
	for {
		len, err := con.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[0:len]))
	}
}

func main() {
	StartServer()
}
