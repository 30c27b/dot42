package main

import (
	"flag"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/30c27b/dot42/cmd/load"

	"github.com/30c27b/dot42/cmd/setup"
	"github.com/30c27b/dot42/internal/config"
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatal("error: this tool only works on macOS darwin")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error: cannot query home directory")
	}

	configPath := flag.String("c", path.Join(homeDir, ".42/configs/config.json"), "/config.json file path")

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
