package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	token := os.Getenv("TOKEN")
	prefix := os.Getenv("PREFIX")

	Token = token
	BotPrefix = prefix

	return nil
}
