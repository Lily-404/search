package search

import (
	"encoding/json"
	"os"
)

var DefaultEngine *string

type Config struct {
	DefaultEngine string `json:"default_engine"`
}

func SetDefaultEngine(s string) {
	DefaultEngine = &s
	saveConfig()
}

func InitConfig() {
	loadConfig()
	if DefaultEngine == nil {
		defaultValue := "google"
		DefaultEngine = &defaultValue
		saveConfig()
	}
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	DefaultEngine = &config.DefaultEngine
}

func saveConfig() {
	config := Config{
		DefaultEngine: *DefaultEngine,
	}

	file, err := os.Create("config.json")
	if err != nil {
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(config)
}
