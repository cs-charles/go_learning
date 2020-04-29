package web_socket

/*import (
	"fmt"
	"net/http"
)

func SocketServer()  {
	fmt.Printf("new connection\n")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
	fmt.Printf(" => closing connection\n")
	ws.Close()

}

func main() {
	http.Handle("/websocket", websocket.Handler(SocketServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
*/
