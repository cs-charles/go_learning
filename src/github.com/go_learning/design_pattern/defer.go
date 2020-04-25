package design_pattern

/**
资源不再被使用时，使用 defer 延迟调用其后的代码，确保资源能够被关闭或返回给连接池。
其次最重要的是从 panic 中恢复程序运行。
*/

/**
关闭文件流
defer f.Close()
*/

/**
解锁
mu.Lock()

defer mu.Unlock()

*/

/**
关闭chan,如果有必要的话

ch := make(chan float64)

defer close(ch)

*/

/**
从panic中恢复

defer func() {

    if err := recover(); err != nil {

        log.Printf(“run time panic: %v”, err)

    }

}
*/

/**
停止一个ticker
tick1 := time.NewTicker(updateInterval)

defer tick1.Stop()
*/

/**
释放一个进程
p, err := os.StartProcess(…, …, …)

defer p.Release()
*/

/**
停止CPU分析并刷新信息
pprof.StartCPUProfile(f)

defer pprof.StopCPUProfile()

*/
