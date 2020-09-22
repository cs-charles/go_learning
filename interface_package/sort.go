package interface_package

import (
	"sort"
)

type IntArray []int

//实现sort.go中的接口，由于p（切片）是引用类型，本本身就是一个指针，所以不用传递指针了
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Convenience wrappers for common cases
func SortInts(a []int) {
	sort.Sort(IntArray(a))
}
func SortStrings(a []string) {
	//说明结构体可以作为方法(参数为某个接口)的参数
	var s sort.Interface = StringArray(a)
	sort.Sort(s)
}
