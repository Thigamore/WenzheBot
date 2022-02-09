package main

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	f, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}

type Bot struct {
	Name string `json: "name"`
}

//Called when a message is created
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Ignore messages by bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(m.Content) == "!wenzhe" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Wenzhe is god")
		if err != nil {
			log.Println("Cannot send message: " + err.Error())
		}
	}
}
