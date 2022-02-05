package bot

import (
	"discord-bot-golang/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Running...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sex_messages := []string{
		"Referência",
		"Isso é uma clara referência a sexo",
		"Estão falando de sexo!? 👀",
		"Tem crianças aqui!",
		"Mas o que é isso? Sexo acidental!",
		"Apenas ocasional sem proteção",
		"Eita porra, tão falando de sexo!",
	}
	if m.Author.ID == BotId {
		return
	}

	content_message := strings.ToLower(m.Content)
	have_sexo := strings.Contains(content_message, "sexo")

	if have_sexo {
		_, _ = s.ChannelMessageSend(m.ChannelID, sex_messages[rand.Intn(7)])
	}
}
