package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Http 是一个比 tcp 更高级的协议，它描述了客户端浏览器如何与网页服务器进行通信。Go 有自己的 net/http 包，我们来看看它。我们从一些简单的示例开始，

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	//另外的输出内容的方式
	//io.WriteString(w, "hello, world!\n")
	/*write := bufio.NewWriter(w)
	write.WriteString()
	write.Flush()*/

	//获取参数
	//方式一 var1是参数名称 未找到返回空字符串
	//req.FormValue("var1")
	//方式二 先执行parseForm，再获取参数。found属性代表是否存在
	//req.ParseForm()
	//var1, found := req.Form["var1"]

}

func StartWebServer() {
	//匹配路径
	http.HandleFunc("/", HelloServer)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
		panic(err)
	}
}

func TestHead() {
	var urls = []string{
		"http://www.baidu.com/",
		"http://golang.org/",
		"http://blog.golang.org/",
	}
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		buf, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(url, ": ", resp.Status, "content:", string(buf))
	}
}

func TestGet() {
	res, err := http.Get("http://www.baidu.com")
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

/**http.Redirect(w ResponseWriter, r *Request, url string, code int)：这个函数会让浏览器重定向到 url（是请求的 url 的相对路径）以及状态码。
http.NotFound(w ResponseWriter, r *Request)：这个函数将返回网页没有找到，HTTP 404 错误。
http.Error(w ResponseWriter, error string, code int)：这个函数返回特定的错误信息和 HTTP 代码。
另 http.Request 对象的一个重要属性 req：req.Method，这是一个包含 GET 或 POST 字符串，用来描述网页是以何种方式被请求的。**/
