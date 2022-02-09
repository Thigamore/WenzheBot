package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "OTQwNjkyNjE5ODcxMjg5MzY0.YgLGRQ.EPUVcGPUEBIvv-kx0eHDmhgtVCU", "Bot Token")
	flag.Parse()
}

func main() {
	//Create discord session for bot with token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("Error creating Discord session: " + err.Error())
	}
	defer dg.Close()

	//Deals with message events
	dg.AddHandler(messageHandler)

	//Only receives message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//Open websocket to Discord
	err = dg.Open()
	if err != nil {
		panic("Error opening connection: " + err.Error())
	}

	//Bot is running
	fmt.Println("Bot is running.\nPress Ctrl+C to exit.")

	//Create an os signal aka interrupt
	sc := make(chan os.Signal, 1)

	//Relays signal to sc
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	//Waits for signal aka a kill
	<-sc

}
