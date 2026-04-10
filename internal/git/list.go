package git

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commit struct {
	Message string
	Hash    string
	Author  string
	Date    string
}

func setHead(dir string) (*git.Repository, *plumbing.Reference, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, nil, err
	}

	head, err := repo.Head()
	if err != nil {
		return nil, nil, err
	}

	return repo, head, nil
}

func listCommit(dir string) ([]Commit, error) {
	var commits []Commit
	repo, head, err := setHead(dir)

	if err != nil {
		return commits, err
	}

	iter, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		return commits, err
	}

	iter.ForEach(func(c *object.Commit) error {
		commits = append(commits, Commit{
			Message: strings.Split(c.Message, "\n")[0],
			Hash:    c.Hash.String()[:7],
			Author:  c.Author.Name,
			Date:    c.Author.When.Format("2006-01-02 15:04:05"),
		})
		return nil
	})

	return commits, nil
}
