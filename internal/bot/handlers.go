package bot

import (
	"fmt"
	"log"
	"mybot/internal/models"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) runHandlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	re := regexp.MustCompile(`^!([a-zA-Z]+)`)
	content := m.Content
	switch re.FindString(content) {
	case b.Config.BotPrefix + "help":
		b.Wg.Add(1)
		go b.helpHandler(s, m)
	case b.Config.BotPrefix + "weather":
		b.Wg.Add(1)
		go b.weatherHandler(s, m, content)
	}

}

func (b *Bot) helpHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer b.Wg.Done()
	if _, err := s.ChannelMessageSend(m.ChannelID, models.HelpText); err != nil {
		log.Println(err)
	}
}

func (b *Bot) weatherHandler(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	defer b.Wg.Done()
	city := strings.TrimSpace(content[8:])

	if city == "" {
		if _, err := s.ChannelMessageSend(m.ChannelID, models.WeatherHelp); err != nil {
			log.Println(err)
		}
		return
	}

	weather, err := b.service.Weather.GetWeatherOfCity(b.Config.WeatherToken, city)
	if err != nil {
		if _, err := s.ChannelMessageSend(m.ChannelID, err.Error()); err != nil {
			log.Println(err)
		}
		return
	}

	if len(weather.Weather) == 0 {
		if _, err := s.ChannelMessageSend(m.ChannelID, "invalid city"); err != nil {
			log.Println(err)
		}
		return
	}

	if _, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("weather in %s city %s: %.2f°C", city, weather.Weather[0].Main, weather.Main.Temperature-273.15)); err != nil {
		log.Println(err)
		return
	}

}
