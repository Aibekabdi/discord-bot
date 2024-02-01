package bot

import (
	"fmt"
	"log"
	"mybot/internal/service"
	"mybot/pkg"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	service *service.Service
	Session *discordgo.Session
	Config  *pkg.Conf
	Wg      sync.WaitGroup
}

func NewBot(conf *pkg.Conf, service *service.Service) *Bot {
	return &Bot{
		Config:  conf,
		service: service,
	}
}

func (b *Bot) Start() {
	// Creating session
	session, err := discordgo.New("Bot " + b.Config.BotToken)
	if err != nil {
		log.Fatalf("Error occured while creating a session : %v\n", err)
	}
	b.Session = session

	b.Session.AddHandler(b.runHandlers)

	// Giving mixed Intets
	b.Session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Starting bot session
	if err = b.Session.Open(); err != nil {
		log.Fatalf("Error occured while opening Discord session: %v\n", err)
	}
	fmt.Println("the bot is online!")

}

// function to close discord session
func (b *Bot) Stop() {
	if b.Session != nil {
		b.Session.Close()
		fmt.Println("Bot is offline!")
	}
}
