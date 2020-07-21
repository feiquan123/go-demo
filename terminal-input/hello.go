package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// bufio reader io
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// input = input[:len(input)-1]
	fmt.Println("name is ", input)
}
