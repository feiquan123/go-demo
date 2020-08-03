package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/feiquan123/go-demo/protobuf-app/proto_gen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	err    error
	logger = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	serverAddress = flag.String("s", "localhost:8888", "server ip address")
	msg           string
	conn          net.Conn
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *serverAddress == "" {
		logger.Panicln("server address is empty")
	}

	// connection
	conn, err = net.Dial("tcp", *serverAddress)
	if err != nil {
		logger.Panicln(err)
	} else {
		logger.Printf("[x] -  connect server[%v] success", *serverAddress)
	}
}

func write(conn net.Conn, data []byte) {
	n, err := conn.Write(data)
	if err != nil {
		logger.Printf("send msg[%d]:%v ,error:%v\n", n, data, err)
	} else {
		logger.Printf("send msg[%d]:%s\n", n, data)
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
			return
		}
	}
}

func handle(conn net.Conn, data []byte) {
	logger.Printf("receive msg[%d]:%s\n", len(data), data)
	msg := &proto_gen.Message{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Printf("receive data[%v]: %s", conn.RemoteAddr(), msg.Message)
	}
}

func main() {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		if err := conn.SetDeadline(time.Now().Add(1e9 * 50)); err != nil {
			logger.Panicln(err)
		}

		fmt.Print("bash :")
		line, err := reader.ReadString('\n')
		if err != nil {
			logger.Panicln(err)
		}
		line = line[:len(line)-1]
		msg = string(line)

		if strings.ToLower(msg) == "q" || strings.ToLower(msg) == "exit" {
			return
		}

		// encode data
		message := &proto_gen.Message{
			Message: string(msg),
			Length:  protoimpl.SizeCache(len(msg)),
		}

		msgData, err := proto.Marshal(message)
		if err != nil {
			logger.Panicln(err)
		}

		// send msg
		write(conn, msgData)
		read(conn, handle)
	}
}
