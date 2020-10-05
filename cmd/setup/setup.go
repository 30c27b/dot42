package setup

import (
	"fmt"

	"github.com/30c27b/dot42/internal/brew"
)

func Run() {
	brew.Install()
	fmt.Println("Done.")
}
