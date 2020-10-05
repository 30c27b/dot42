package brew

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/30c27b/dot42/internal/config"
	"github.com/30c27b/dot42/internal/utils"
)

type brewCtx struct {
	homeDir      string
	fullPath     string
	repoFullPath string
	caskPath     string
	brewExec     string
}

// Install checks if brew is installed and installs it if it is not the case
func Install() {
	ctx := newBrewCtx()
	if localBrewIsOnPath(ctx) {
		fmt.Println("warning: brew is already installed in your home directory, you may want to update your path to ignore it.")
	}
	createDir(ctx)
	if !utils.IsEmpty(ctx.repoFullPath) {
		log.Fatal("the directory", ctx.repoFullPath, "should be empty.")
	}
	localInstall(ctx)
}

func newBrewCtx() *brewCtx {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error: cannot query home directory")
	}
	fullPath := path.Join(homeDir, config.Cfg.BrewPath)
	repoFullPath := path.Join(fullPath, config.Cfg.BrewRepository)
	caskPath := path.Join(homeDir, config.Cfg.BrewCaskPath)
	brewExec := path.Join(fullPath, "/bin/brew")
	return &brewCtx{homeDir, fullPath, repoFullPath, caskPath, brewExec}
}

func localBrewIsOnPath(ctx *brewCtx) bool {
	p, err := exec.LookPath("brew")
	if err != nil {
		return false
	}
	if !strings.HasPrefix(p, ctx.homeDir) {
		return false
	}
	return true
}

func createDir(ctx *brewCtx) {
	if err := os.MkdirAll(ctx.repoFullPath, 0755); err != nil {
		log.Fatal("error: the directory", ctx.repoFullPath, "could not be created.")
	}
	if err := os.Mkdir(path.Join(ctx.fullPath, "bin"), 0755); err != nil {
		log.Fatal("error: the directory", ctx.fullPath+"/bin", " could not be created.")
	}
	if err := os.Mkdir(ctx.repoFullPath, 0755); err != nil {
		log.Fatal("error: the directory", ctx.repoFullPath, " could not be created.")
	}
	if err := os.Mkdir(ctx.caskPath, 744); err != nil {
		log.Fatal("error: the directory", ctx.caskPath, "could not be created.")
	}
}

// LocalInstall clones, installs and updates brew
func localInstall(ctx *brewCtx) {
	if _, err := exec.Command("git", "clone", config.Cfg.BrewRepository, ctx.repoFullPath).Output(); err != nil {
		log.Fatal("error: brew could not be cloned:", err)
	}
	if _, err := exec.Command("ln", "-sf", path.Join(ctx.repoFullPath, "bin/brew"), ctx.brewExec).Output(); err != nil {
		log.Fatal("error: brew could not be linked to", ctx.brewExec, ":", err)
	}
	if _, err := exec.Command(ctx.brewExec, "update", "--force").Output(); err != nil {
		log.Fatal("error: Brew could not be updated", err)
	}
}
