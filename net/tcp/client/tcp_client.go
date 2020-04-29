package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func startClient() {
	con, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Dialing", err.Error())
		return
	}
	read := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := read.ReadString('\n')
	clientName = strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := read.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		// fmt.Printf("input:--s%--", input)
		// fmt.Printf("trimmedInput:--s%--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, _ = con.Write([]byte(clientName + " says: " + trimmedInput))
	}
}

func main() {
	startClient()
}
