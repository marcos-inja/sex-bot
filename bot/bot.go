package bot

import (
	"discord-bot-golang/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
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
	goBot.AddHandler(jokeResp)
	goBot.AddHandler(setPresence)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Running...")
}

func setPresence(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0,"videos no xvideos!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sex_messages := [8]string{
		"Referência",
		"Isso é uma clara referência a sexo",
		"Estão falando de sexo!? 👀",
		"Tem crianças aqui!",
		"Mas o que é isso? Sexo acidental!",
		"Apenas ocasional sem proteção",
		"Eita porra, tão falando de sexo!",
		"Acho que isso é uma referência a sexo",
	}
	if m.Author.ID == BotId {
		return
	}

	content_message := strings.ToLower(m.Content)
	have_sexo := strings.Contains(content_message, "sexo")

	if have_sexo {
		_, _ = s.ChannelMessageSend(m.ChannelID, sex_messages[rand.Intn(8)])
	}
}

func jokeResp(s *discordgo.Session, m *discordgo.MessageCreate) {

	file, err := ioutil.ReadFile("./jokes/jokes.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	var arr []string
	json.Unmarshal([]byte(file), &arr)
	
	if m.Content == config.BotPrefix + "piada"{
		length_arr := len(arr)
		_, _ = s.ChannelMessageSend(m.ChannelID, arr[rand.Intn(length_arr)])
	}
}
