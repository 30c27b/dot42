package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Cfg stores information about the config file
var Cfg Config

// Config stores parameters from the config.json file
type Config struct {
}

// Process transforms a config file into a Config object
func Process(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}
