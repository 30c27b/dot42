package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/30c27b/dot42/cmd/load"

	"github.com/30c27b/dot42/cmd/setup"
	"github.com/30c27b/dot42/internal/config"
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatal("error: this tool only works on macOS darwin")
	}

	configPath := flag.String("c", "configs/config.json", "/config.json file path")

	flag.Parse()

	config.Process(*configPath)

	if len(os.Args) == 1 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setup":
		setup.Run()
	case "load":
		load.Run()
	default:
		os.Exit(1)
	}
}
