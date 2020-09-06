package main

import (
	"bufio"
	"fmt"
	"os"

	"go.nanomsg.org/mangos/v3/protocol/req"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

const listenURL = "ipc:///tmp/test.ipc"

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func main() {
	sock, err := req.NewSocket()
	if err != nil {
		die("failed to create socket: %v", err)
	}

	if err = sock.Dial(listenURL); err != nil {
		die("failed to dial %v: %v", listenURL, err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your name: ")
		if ok := scanner.Scan(); !ok {
			continue
		}

		if err := sock.Send([]byte(scanner.Text())); err != nil {
			fmt.Println("failed to send message:", err)
			continue
		}

		msg, err := sock.Recv()
		if err != nil {
			fmt.Println("failed to receive message:", err)
			continue
		}

		fmt.Println(string(msg))
	}

}
