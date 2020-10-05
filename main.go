package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/30c27b/dot42/internal/brew"
	"github.com/30c27b/dot42/internal/config"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("this tool only works on macOS darwin")
	}

	configPath := flag.String("configPath", "configs/config.json", "config file path")
	config.Process(*configPath)

	brew.Install()
}
