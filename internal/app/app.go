package app

import (
	"fmt"
	"log"
	"mybot/pkg"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	botID string
)

func Run(conf *pkg.Conf) {
	session, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		log.Fatalf("error occured while creates a new session by token :%v\n", err)
	}
	// user, err := session.User("@me")
	// if err != nil {
	// 	log.Println("lol", err)
	// 	return
	// }
	// botID = user.ID
	fmt.Println(botID)
	session.AddHandler(messageHandler)

	if err = session.Open(); err != nil {
		log.Println("lol", err)
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("Received message: %v \n", m.Content)
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			log.Println(err)
		}

	}
}
