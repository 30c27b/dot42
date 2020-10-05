package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Cfg stores information about the config file
var Cfg Config

// Config stores parameters from the config.json file
type Config struct {
	BrewName       string
	BrewPath       string
	BrewCaskPath   string
	BrewRepository string
}

// Process transforms a config file into a Config object
func Process(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("error: could not open config file at", path)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Cfg); err != nil {
		fmt.Println("error: config file could not be parsed:", err)
	}
}
