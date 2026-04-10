package git

import (
	"errors"
	"fmt"

	goGit "github.com/go-git/go-git/v5"
	"github.com/lukas-sgx/seniorlens/internal/ui"
)

func errorRepoNotFound(err error) bool {
	if errors.Is(err, goGit.ErrRepositoryNotExists) {
		fmt.Println(ui.SetYellow("WARNING: No git repository found in this directory."))
		return true
	}
	return false
}

func errorNoCommits(err error) bool {
	if errors.Is(err, goGit.ErrEmptyCommit) {
		fmt.Println(ui.SetYellow("WARNING: No commits found in this repository."))
		return true
	}
	return false
}