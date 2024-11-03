package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	AirconIP string `json:"aircon_ip"`
}

func GetConfig() (Config, error) {
	var config Config
	configPath := fmt.Sprintf("%s/.config/aircon/aircon.json", os.Getenv("HOME"))
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return config, errors.New("Could not find a config at ~/.config/aircon/aircon.json. Please check the README.")
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		return config, errors.New(fmt.Sprintf("%v\nPlease check the README.", err))
	}

	return config, nil
}
