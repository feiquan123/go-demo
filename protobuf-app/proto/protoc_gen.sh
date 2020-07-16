mkdir ../proto_gen
protoc --proto_path=./ --go_out=../proto_gen/ *.proto