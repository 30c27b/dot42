package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/30c27b/dot42/internal/config"
	"github.com/30c27b/dot42/internal/pref"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("this tool only works on macOS darwin")
	}

	configPath := flag.String("configPath", "config/config.json", "config file path")
	config.Process(*configPath)

	p := pref.New("Apple Global Domain", "AppleInterfaceStyle", "string", "Dark")
	v := p.Read()
	fmt.Println(v)
}
