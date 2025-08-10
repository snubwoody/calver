package main

import (
	"fmt"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	tags, err := repo.Tags()
	if err != nil {
		panic(err)
	}

	err = tags.ForEach(func(tag *plumbing.Reference) error {
		fmt.Printf("Tag: %v", tag.String())
		return nil
	})

	if err != nil {
		panic(err)
	}
}
