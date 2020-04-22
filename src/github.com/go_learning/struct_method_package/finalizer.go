package struct_method_package

import (
	"fmt"
	"runtime"
)

func Finalizer() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	//上面的程序会给出已分配内存的总量，单位是 Kb。
	fmt.Printf("%d Kb\n", m.Alloc/1024)
	i := 1000
	//在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，即使程序正常结束或者发生错误。
	runtime.SetFinalizer(i, func(obj *int) {
		fmt.Printf("i value is %v\n", 123)
	})

}
