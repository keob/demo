package main

import (
	"fmt"
	"net"

	"github.com/keob/demo/socket/tcp/proto"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:30000")
	if err != nil {
		fmt.Println("dial failed, err: ", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode message failed, err: ", err)
			return
		}

		conn.Write(data)
	}
}
