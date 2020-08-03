package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"github.com/feiquan123/go-demo/protobuf-app/proto_gen"
	"github.com/feiquan123/go-demo/protobuf-app/utils/bash"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	logger = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	port = flag.String("p", "8888", "server port")

	recData  []byte
	sendData []byte
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *port == "" {
		logger.Panicln("server port is empty")
	}
}

func read(conn net.Conn, handle func(conn net.Conn, data []byte)) {
	var recdata = make([]byte, 0)
	defer func() {
		recdata = make([]byte, 0)
	}()

	buf := make([]byte, 4096)
	for {
		l, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			logger.Println(err)
		}
		recdata = append(recdata, buf[:l]...)
		if l >= len(buf) {
			continue
		}

		if handle != nil {
			handle(conn, recdata)
		}
	}
}

func write(conn net.Conn, data []byte) {
	n, err := conn.Write(data)
	if err != nil {
		logger.Printf("send msg[%d]:%v ,error:%v\n", n, data, err)
	} else {
		// logger.Printf("send msg[%d]:%s\n", n, data)
	}
}

func handle(conn net.Conn, data []byte) {
	msg := &proto_gen.Message{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Printf("receive data[%v]: %s", conn.RemoteAddr(), msg.Message)
	}

	// run cmd
	send := ""
	sendMsg, err := bash.Cmd(msg.Message)
	if err != nil {
		send = err.Error()
	} else {
		send = sendMsg.String()
	}

	// encode data
	message := &proto_gen.Message{
		Message: send,
		Length:  protoimpl.SizeCache(len(msg.Message)),
	}

	msgData, err := proto.Marshal(message)
	if err != nil {
		logger.Panicln(err)
	}
	write(conn, msgData)
}

func main() {
	address := "localhost:" + *port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Panicln(err)
	}
	logger.Printf("[x] - start server , listen:%s", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Println(err)
		}
		// ansync handle
		go read(conn, handle)
	}
}
