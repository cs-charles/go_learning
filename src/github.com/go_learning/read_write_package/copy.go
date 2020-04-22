package read_write_package

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//整个文件拷贝
func CopyFile1() {
	inputFile := "./files/products.txt"
	outputFile := "./files/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

func CopyFile2(src, dest string) {
	srcFile, openFileErr := os.Open(src)
	if openFileErr != nil {
		panic(openFileErr)
	}
	defer srcFile.Close()

	destFile, writeFileErr := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if writeFileErr != nil {
		panic(writeFileErr)
	}
	defer destFile.Close()
	io.Copy(destFile, srcFile)
}
