package main

import (
	"fmt"
	"os"

	"go.nanomsg.org/mangos/v3/protocol/rep"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

const listenURL = "ipc:///tmp/test.ipc"

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func main() {

	sock, err := rep.NewSocket()
	if err != nil {
		die("failed to create socket: %v", err)
	}

	if err = sock.Listen(listenURL); err != nil {
		die("failed to listen on socket: %v", err)
	}

	fmt.Println("Listen for message on pipe:", listenURL)
	for {
		fmt.Println("Waiting for connection...")
		msg, err := sock.Recv()
		if err != nil {
			fmt.Println("error while receiving msg:", err)
			continue
		}
		fmt.Println("Received message:", string(msg))
		//time.Sleep(15 * time.Second)

		if err = sock.Send([]byte(fmt.Sprintf("Hello %v, how are you today?", string(msg)))); err != nil {
			fmt.Println("failed to send message:", err)
			continue
		}
		fmt.Println("Successfully send message...")
	}

}
