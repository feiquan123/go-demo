package main

import (
	"errors"
	"fmt"
	"time"
)

type PlayerParam string

type Player struct {
	Name string
}

func NewPlayer(name PlayerParam) (Player, func(), error) {
	clearup := func() {
		fmt.Println("player clearup!")
	}
	if time.Now().Unix()%2 == 0 {
		return Player{}, clearup, errors.New("palyer dead")
	}

	return Player{
		Name: string(name),
	}, clearup, nil
}
