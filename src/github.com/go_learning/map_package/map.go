package map_package

import "fmt"

//定义 var map1 map[keytype]valuetype
//未初始化时，map1是nil
//key 可以是任意可以用 == 或者 != 操作符比较的类型，
////比如 string、int、float。所以数组、切片和结构体不能作为 key
//(译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的）
//值可以是任意类型的,比如方法func()
var map1 map[string]int
var map2 = map[string]int{"key1": 1, "key2": 2}

//map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。
//通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；
//所以如果你很在乎性能的话还是建议用切片来解决问题。
var map3 = make(map[int]string, 10)

//不要使用 new，永远用 make 来构造 map

var map4 = map[int][]int{1: {1, 3, 5}, 2: {2, 4, 6}}

//判断map是否存在某个key
func IsPresent(key int) bool {
	_, isPresent := map4[key]
	return isPresent
}

func DeleteByKey(key string) {
	delete(map2, key)
}

func InterMap() {
	for key, value := range map4 {
		fmt.Printf("key is %d, value is %v \n", key, value)
	}
}

//map类型的切片
var slice1 = []map[string]string{{"1": "1", "2": "2"}, {"a": "a", "b": "b"}}
var slice2 = make([]map[string]string, 10, 10)

func InitSlice() {
	slice2[0] = make(map[string]string, 10)
	slice2[0]["key1"] = "value1"
}

//map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序（详见第 8.3 节）。
//如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序
//如果你想要一个排序的列表你最好使用结构体切片,比如下面这个
//type name struct {
//    key string
//    value int
//}
