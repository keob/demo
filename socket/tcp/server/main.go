package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/keob/demo/socket/tcp/proto"
)

func process(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode message failed, err: ", err)
			return
		}
		fmt.Println("Receive client message: ", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:30000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}
}
