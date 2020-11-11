package main

import (
	"fmt"
)

var me string

func main() {
	fmt.Print(me)
	fmt.Print(string(rune(96)))
	fmt.Print(me)
	fmt.Print(string(rune(96)) + "\n}")
}

func init() {
	me = `package main

import (
	"fmt"
)

var me string

func main() {
	fmt.Print(me)
	fmt.Print(string(rune(96)))
	fmt.Print(me)
	fmt.Print(string(rune(96))+"\n}")
}

func init() {
	me = `
}