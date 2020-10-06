package load

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/30c27b/dot42/internal/config"
)

func Run() {
	if config.Cfg.DarkTheme {
		setDarkTheme()
	}
	for _, pref := range config.Cfg.Defaults {
		pref.Write()
	}
	fmt.Println("configuration loaded.")
}

func setDarkTheme() {
	if _, err := exec.Command("osascript", "-e", "tell app \"System Events\" to tell appearance preferences to set dark mode to true").Output(); err != nil {
		log.Fatal("error: could not set dark theme.")
	}
}
