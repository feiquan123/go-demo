package main

import "fmt"

type EndingB struct{
	Player Player
	Moster Moster
}

func (e EndingB)Appear(){
	fmt.Printf("%s defeats %s, but becom monster, world darker!\n",e.Player.Name,e.Moster.Name)
}