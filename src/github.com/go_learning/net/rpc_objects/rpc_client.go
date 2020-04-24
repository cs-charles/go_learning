package rpc_objects

import (
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func StartRpcClient() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// 同步调用
	args := &Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	//异步调用
	/**
	call1 := client.Go("Args.Multiply", args, &reply, nil)

	replyCall := <- call1.Done
	*/
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
	//https://learnku.com/docs/the-way-to-go/1510-uses-netchan-across-network-to-implement-message-delivery/3712
}
