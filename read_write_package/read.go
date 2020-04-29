package read_write_package

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
//Scanf 与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
//Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，
//与 Scanf 相同。如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误。

//比较：
//scan函数会识别空格左右的内容，哪怕换行符号存在也不会影响scan对内容的获取。
//scanln函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容。

//Scan、Scanf和Scanln: 从标准输入os.Stdin读取文本(从终端获取数据)
//Fscan、Fscanf、Fscanln: 从指定的io.Reader接口读取文本(通用)
//Sscan、Sscanf、Sscanln: 从一个参数字符串读取文本(从字符串string获取数据)
func ReadFromStdIn() {
	fmt.Println("scan")
	//scan 从标准输入流中读取
	var str1, str2 string
	n1, _ := fmt.Scan(&str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n1, str1, str2)
	//sscan 从字符串中读取 sscan会把换行符当做空格处理
	input := "56.12\n5212 Go"
	var f1 float32
	var i1 int
	var str3 string
	n2, _ := fmt.Sscan(input, &f1, &i1, &str3)
	fmt.Printf("读取了 %v 个变量从字符串,分别是f1= %v,i1=%d,str3=%s\n", n2, f1, i1, str3)
	n3, _ := fmt.Fscan(os.Stdin, &str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n3, str1, str2)

	fmt.Println("scanln")
	//scanln
	n4, _ := fmt.Scanln(&str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n4, str1, str2)
	input = "56.12 5212 Go"
	n5, _ := fmt.Sscan(input, &f1, &i1, &str3)
	fmt.Printf("读取了 %v 个变量从字符串,分别是f1= %v,i1=%d,str3=%s\n", n5, f1, i1, str3)
	n6, _ := fmt.Fscanln(os.Stdin, &str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n6, str1, str2)

	//scanf
	fmt.Println("scanf")
	n7, _ := fmt.Scanf("str1=%s str2=%s", &str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n7, str1, str2)
	input = "56.12 / 5212 / Go"
	format := "%f / %d / %s"
	n8, _ := fmt.Sscanf(input, format, &f1, &i1, &str3)
	fmt.Printf("读取了 %v 个变量从字符串,分别是f1= %v,i1=%d,str3=%s\n", n8, f1, i1, str3)
	n9, _ := fmt.Fscanf(os.Stdin, "str1=%s,str2=%s", &str1, &str2)
	fmt.Printf("读取了 %v 个变量从控制台,分别是str1= %s,str2=%s\n", n9, str1, str2)

	//利用bufio读取
	fmt.Println("bufio")
	var read *bufio.Reader = bufio.NewReader(os.Stdin)
	inputString, err := read.ReadString('\n')
	if err != nil {
		fmt.Println("输入出错，没有换行符")
	} else {
		fmt.Printf("输入内容是%s", inputString)
	}
}

//go中的File类型实现了Read接口
func ReadFile1() {
	dir, _ := os.Getwd()
	//输出go_work_space的路径
	fmt.Printf("curror path is %s\n", dir)
	file, err := os.Open("./files/readFile.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n"+
			"Does the file exist?\n"+
			"Have you got acces to it?\n"+
			"error is %v", err)
		return // exit the function on error
	}
	defer file.Close()
	//file可以直接读，下面是采用了缓冲区
	readFile := bufio.NewReader(file)
	for {
		inputString, inputErr := readFile.ReadString('\n')
		fmt.Printf("The input was: %s\n", inputString)
		if inputErr == io.EOF {
			return
		}
	}

}

//带缓冲的读取
func ReadFile2() {
	file, _ := os.Open("./files/readFile.txt")
	defer file.Close()
	buf := make([]byte, 1024)
	read := bufio.NewReader(file)
	for {
		n, err := read.Read(buf)
		fmt.Printf("The input number was: %d\n", n)
		fmt.Printf("The input was: %s\n", string(buf))
		if err == io.EOF {
			return
		}
	}

}

//按列读取
func ReadCol() {
	inputFile := "./files/inputFile.txt"
	file, err1 := os.Open(inputFile)
	if err1 != nil {
		panic(err1)
	}
	defer file.Close()
	var col1, col2, col3 string
	var slice1, slice2, slice3 []string
	for {
		_, err2 := fmt.Fscanln(file, &col1, &col2, &col3)
		if err2 != nil {
			break
		}
		slice1 = append(slice1, col1)
		slice2 = append(slice2, col2)
		slice3 = append(slice3, col3)
	}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
