package main

import (
	"discord-ping-bot/bot"
	"discord-ping-bot/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
