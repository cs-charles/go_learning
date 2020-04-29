package netchan

//netchan 跨网络实现消息传递
/**
一个与 rpc 密切相关的在网络上使用 通道 的技术。在 [14 章]() 中，通道作为一个本地通道被使用，它们只存在于执行它们的机器的内存空间中。
netchan 包实现了类型安全的网络通道： 它允许通道的两端出现在通过网络连接的不同计算机上。
一个出口按照名称发布 一个（组） 通道 。一个入口去连接出口的机器，并按照名称输入到 通道 。网络 通道 不是同步的，
它们就像是缓冲 通道
*/

/*func StartNetChanIn()  {
	imp, err := netchan.NewImporter("tcp", "netchanserver.mydomain.com:1234")
	if err != nil {
		log.Fatalf("Error making Importer: %v", err)
	}
	ch := make(chan myType)
	err = imp.Import("sendmyType", ch, netchan.Receive)
	if err != nil {
		log.Fatalf("Receive Error: %v", err)
	}
}*/
