package main

import (
	"fmt"
	"log"
	"mybot/internal/bot"
	"mybot/internal/models"
	"mybot/internal/service"
	"mybot/pkg"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Saving logs into file
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("cannot create log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
	// Getting configurations from configs.json
	conf, err := pkg.NewConfig(models.ConfigPath)
	if err != nil {
		log.Fatalf("error occured while getting configs from file: %v", err)
	}

	services := service.NewService()
	// Creating a new struct with Discord bot session
	b := bot.NewBot(conf, services)
	// Running bot
	b.Start()

	// making our non stopable while it is not Fatalled or terminal stopping
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc

	timeout := time.After(10 * time.Second)

	select {
	case <-timeout:
		fmt.Println("Timeout exceeded. Exiting program.")
	case <-waitForCompletion(b):
		fmt.Println("All goroutines completed. Exiting program.")
	}
	// Closing our session
	b.Stop()
}

func waitForCompletion(b *bot.Bot) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		b.Wg.Wait()
		close(ch)
	}()
	return ch
}
