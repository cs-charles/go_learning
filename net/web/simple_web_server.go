package web

import (
	"io"
	"log"
	"net/http"
)

const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

type HandleFnc func(http.ResponseWriter, *http.Request)

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	//http.DetectContentType() 可以检测内容的格式
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data */
		//request.ParseForm();
		//io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in"))
	}
}

func StartSimpleWebServer() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

/**
 处理panic
当 web 服务器发生一个恐慌（ panic ）时，我们的 web 服务器就会终止。
这样非常的糟糕：一个 web 服务必须是一个健壮的程序，能够处理可能会出现的问题。

*/
func logPanics(function HandleFnc) HandleFnc {

	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {

			if x := recover(); x != nil {

				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)

				// 下面一行代码是译者添加，默认出现 panic 只会记录日志，页面就是一个无任何输出的白页面，
				// 可以给页面一个错误信息，如下面的示例返回了一个 500
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}

		}()

		function(writer, request)
	}

}
