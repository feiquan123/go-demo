package main

import (
	"log"
	"time"
)

func main() {
	mission, clearup, err := InitMission("jdk", "jek", time.Now())
	if err != nil {
		log.Println(err)
	}

	mission.Start()
	clearup()
}
