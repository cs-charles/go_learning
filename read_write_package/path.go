package read_write_package

import (
	"fmt"
	"os"
	"path"
)

func FilePath() {
	dir, _ := os.Getwd()
	//获取路径中最后一个元素名称
	lastPathEle := path.Base(dir)
	fmt.Println(lastPathEle)
	filePath1 := path.Join(dir, "files")
	fmt.Println(filePath1)
}
