package arr_with_slice_package

import "fmt"

//copy函数，将一个切面中的内容复制到另外一个切面中去
//并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。
func Copy() {
	var dest = make([]int, 3, 10)
	var src = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	count := copy(dest, src) //return 3，返回src与dest中长度最小的那个
	fmt.Printf("copy count is %d,dest value is %v \n", count, dest)
}

//append函数，切片追加新元素
//；追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。
// 如果想将y切片追加到x切面后面， x = append(x,y...)
//返回来的切片长度等于两个之和，切片容量等于最大的切片容量或者新的切片容量
func Append() {
	var y = []int{12, 13, 14}
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	x = append(x, y...)
	fmt.Printf("x value is %v \n", x)
}

//其他用法
//切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
//为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
// 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
//在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
//在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
//取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1:], a[:len(a)-1]

func GrowthCap(slice []int, factor int) {
	fmt.Printf("old slice length is %d \n", len(slice))
	//新的长度
	newLen := len(slice) * factor
	//旧的长度
	oldLen := len(slice)
	newSpace := make([]int, newLen-oldLen, newLen-oldLen)
	slice = append(slice, newSpace...)
	fmt.Printf("new slice length is %d \n", len(slice))
}

func Filter(slice []int, f func(i int) bool) []int {
	newSlice := make([]int, len(slice), len(slice))
	for _, val := range slice {
		if f(val) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}
