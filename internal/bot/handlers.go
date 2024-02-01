package bot

import (
	"fmt"
	"log"
	"mybot/internal/models"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) runHandlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// parsing command by regex
	re := regexp.MustCompile(`^!([a-zA-Z]+)`)
	content := m.Content
	switch re.FindString(content) {
	case b.Config.BotPrefix + "help":
		b.Wg.Add(1)
		go b.helpHandler(s, m)
	case b.Config.BotPrefix + "weather":
		b.Wg.Add(1)
		go b.weatherHandler(s, m, content)
	case b.Config.BotPrefix + "remindme":
		b.Wg.Add(1)
		go b.remindmeHandler(s, m, content)
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
	// getting city
	city := strings.TrimSpace(strings.TrimPrefix(m.Content, "!weather"))
	if city == "" {
		if _, err := s.ChannelMessageSend(m.ChannelID, models.WeatherHelp); err != nil {
			log.Println(err)
		}
		return
	}
	// getting weather by openweathermap
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

func (b *Bot) remindmeHandler(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	defer b.Wg.Done()
	// getting time from content
	content = content[9:]
	args := strings.Fields(content)
	if len(args) < 2 {
		if _, err := s.ChannelMessageSend(m.ChannelID, models.ReminderHelp); err != nil {
			log.Println(err)
		}
		return
	}

	timeStr := args[len(args)-1]
	// parsing string time to time
	remindTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		if _, err := s.ChannelMessageSend(m.ChannelID, "not correct time format"); err != nil {
			log.Println(err)
		}
		return
	}
	now := time.Now()
	remindTime = time.Date(now.Year(), now.Month(), now.Day(), remindTime.Hour(), remindTime.Minute(), 0, 0, now.Location())

	// gettinп time when it should remind
	durationUntilRemind := time.Until(remindTime)
	if durationUntilRemind <= 0 {
		if _, err := s.ChannelMessageSend(m.ChannelID, "You can not remind in the past time"); err != nil {
			log.Println(err)
		}
		return
	}
	// make work new goroutine wich will wait untill remind time comes
	timer := time.NewTimer(durationUntilRemind)
	go func() {
		<-timer.C
		if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("i am remiding you!!! %s : %s", m.Author.Mention(), strings.Join(args[:len(args)-1], " "))); err != nil {
			log.Println(err)
			return
		}
	}()

	if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ok, i will remind you at %s.", remindTime.Format("15:04"))); err != nil {
		log.Println(err)
		return
	}
}
