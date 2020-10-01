package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/30c27b/dot42/internal/config"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("this tool only works on macOS darwin")
		os.Exit(2)
	}

	configPath := flag.String("configPath", "config/config.json", "config file path")
	config.Process(*configPath)
}
