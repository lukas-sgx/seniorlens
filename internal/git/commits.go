package git

import (
	"fmt"
)

func GetCommits(dir string) []Commit {
	commits, err := listCommit(dir)

	if err != nil {
		if !errorRepoNotFound(err) {
			fmt.Println(err)
		}
		return nil
	}

	if len(commits) == 0 {
		if !errorNoCommits(err) {
			fmt.Println(err)
		}
		return nil
	}
	return commits
}


func LastCommit(dir string) *Commit {
	commits := GetCommits(dir)

	if commits == nil {
		return nil
	}

	return &commits[0]
}
