package main

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/agungfir98/mini-redis/handler"
	"github.com/agungfir98/mini-redis/proto"
)

func main() {
	fmt.Println("TCP running on PORT: 6380")
	ln, err := net.Listen("tcp", ":6380")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		resp := proto.NewResp(conn)
		writer := proto.NewWriter(conn)

		msg, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			msg := proto.RespMessage{Typ: "error", Error: fmt.Sprintf("ERR %v\n", err)}
			writer.Write(msg)
			continue
		}

		cmd := strings.ToUpper(msg.Array[0].String)
		args := msg.Array[1:]

		if msg.Typ != "array" {
			msg := proto.RespMessage{Typ: "error", Error: "Type error, expected array"}
			writer.Write(msg)
			continue
		}

		handler, ok := handler.Message[cmd]
		if !ok {
			msg := proto.RespMessage{Typ: "error", Error: fmt.Sprintf("ERR unknown command '%v', with args beginning with:", cmd)}
			writer.Write(msg)
			continue
		}

		res := handler(args)
		writer.Write(res)
	}
}
