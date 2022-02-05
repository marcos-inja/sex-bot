package main

import (
	"discord-bot-golang/bot"
	"discord-bot-golang/config"
	"fmt"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
