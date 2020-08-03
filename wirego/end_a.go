package main

import "fmt"

type EndingA struct {
	Player Player
	Moster Moster
}


func (e EndingA) Appear(){
	fmt.Printf("%s defeats %s, world peace!\n",e.Player.Name,e.Moster.Name)
}

