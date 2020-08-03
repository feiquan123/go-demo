package main

import (
	"log"
	"os"

	"github.com/feiquan123/go-demo/protobuf-app/proto_gen"

	"google.golang.org/protobuf/proto"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	p := &proto_gen.Person{
		Name: "Jack",
		Age:  10,
		From: "China",
	}
	logger.Println("", p)

	// 序列化
	dataMarshal, err := proto.Marshal(p)
	if err != nil {
		logger.Panic(err)
	}
	logger.Println("编码数据:", dataMarshal)

	// 反序列化
	entity := proto_gen.Person{}
	err = proto.Unmarshal(dataMarshal, &entity)
	if err != nil {
		logger.Panic(err)
	}
	logger.Printf("解码数据: name:%s age:%d from:%s", entity.Name, entity.Age, entity.From)
}
