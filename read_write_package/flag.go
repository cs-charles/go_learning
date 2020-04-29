package read_write_package

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var EnableNewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func PrintByFlag() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	//flag.VisitAll()
	//flag.VisitAll(fn func(*Flag)) 是另一个有用的功能：按照字典顺序遍历 flag，并且对每个标签调用 fn
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *EnableNewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

func cat1(read *bufio.Reader) {
	for {
		buf, err := read.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

//利用切片读取读取文件
func cat2(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func TestCat() {
	var hasCat = flag.Bool("cat", false, "review file")
	fmt.Println(*hasCat)
	//打印flag注册的所有可选
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 || !*hasCat {
		cat1(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat1(bufio.NewReader(f))
	}
}

//防止切片占据的数组无法被垃圾回收，可以采用copy函数复制切片
