package read_write_package

import (
	"bufio"
	"fmt"
	"os"
)

//如果想写入string到Writer接口可以使用io.WriterString(*Writer,string)

//不使用缓冲区
func WriteFile1() {
	file, _ := os.OpenFile("./files/outputFile1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	file.Write([]byte("哈哈哈哈哈哈"))
	file.WriteString("嘻嘻嘻嘻嘻")
}

//使用缓冲区
func WriteFile2() {
	//os.O_RDONLY：只读
	//os.O_WRONLY：只写
	//os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	//os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0。
	//0666中的0代表0进制
	outputFile, outputError := os.OpenFile("./files/outputFile2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
