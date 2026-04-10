package cli

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/lukas-sgx/seniorlens/internal/git"
	"github.com/lukas-sgx/seniorlens/internal/ui"
)

func getDirName(dir string) string {
	cleanPath := filepath.Clean(dir)

	absolutePath, err := filepath.Abs(cleanPath)
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Base(absolutePath)
}

func commitIdentify() {
	last := git.LastCommit(directory)
	var commit string

	if last == nil {
		return
	}

	commit = ui.SetGreen("Commit: " + ui.SetUnderline(last.Hash)) +
	ui.SetGreen(" (" + last.Author + " - " + last.Date +") - '" + last.Message + "'")
	fmt.Println(commit)
}

func actualProject() {
	project := getDirName(directory)
	scanProject := ui.SetGreen("> ") + "Scanning Project: `" + project + "`..."

	fmt.Println(scanProject)
}

func initApp() {
	version := "0.1.0"
	analysisVersion := "0.0.1"

	header(version, analysisVersion)
	actualProject()
	commitIdentify()
}