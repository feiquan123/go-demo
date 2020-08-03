package main

import (
	"fmt"
	"time"
)

type Win struct{
	Time time.Time
	Result bool
}

type Mission struct {
	Player Player
	Moster Moster
	Win  Win
}

func NewMission(p Player, m Moster , t time.Time) Mission {
	return Mission{p, m, Win{t,time.Now().Unix()%2==0}}
}

func NewMission2(p Player,m Moster,  t time.Time) *Mission{
	return &Mission{p,m,Win{t,time.Now().Unix()%2==0}}
}

func (m Mission)Start(){
	fmt.Printf("%s defeats %s, world peace! time:%s result:%v\n",m.Player.Name,m.Moster.Name,m.Win.Time,m.Win.Result)
}