package pkg

import (
	"encoding/json"
	"os"
)

type Conf struct {
	BotToken string `json:"botToken"`
	BotPrefix string `json:"botPrefix"`
}

func NewConfig(path string) (*Conf, error) {
	var newConfig Conf
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(file).Decode(&newConfig); err != nil {
		return nil, err
	}
	return &newConfig, nil
}
