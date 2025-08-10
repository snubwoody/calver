package main

import (
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestVersionExists(t *testing.T) {
	dir, err := os.MkdirTemp("", "test-output-")
	assert.Nil(t, err)
	defer os.RemoveAll(dir)

	repo, err := git.PlainInit(dir, false)
	assert.Nil(t, err)

	path := filepath.Join(dir, "test.txt")
	_, err = os.Create(path)
	assert.Nil(t, err)

	worktree, err := repo.Worktree()
	assert.Nil(t, err)

	_, err = worktree.Add(path)
	assert.Nil(t, err)

	_, err = worktree.Commit("test commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "automated-test",
			Email: "test-worker@email.com",
			When:  time.Now(),
		},
	})
	assert.Nil(t, err)

	ref, err := repo.Head()
	assert.Nil(t, err)

	_, err = repo.CreateTag("v0.1.0", ref.Hash(), nil)
	assert.Nil(t, err)

	exists, err := VersionExists(repo, "v0.1.0")
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestVersionDoesntExist(t *testing.T) {
	dir, err := os.MkdirTemp("", "test-output-")
	assert.Nil(t, err)
	defer os.RemoveAll(dir)

	repo, err := git.PlainInit(dir, false)
	assert.Nil(t, err)

	path := filepath.Join(dir, "test.txt")
	_, err = os.Create(path)
	assert.Nil(t, err)

	worktree, err := repo.Worktree()
	assert.Nil(t, err)

	_, err = worktree.Add(path)
	assert.Nil(t, err)

	_, err = worktree.Commit("test commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "automated-test",
			Email: "test-worker@email.com",
			When:  time.Now(),
		},
	})
	assert.Nil(t, err)

	ref, err := repo.Head()
	assert.Nil(t, err)

	_, err = repo.CreateTag("v0.1.0", ref.Hash(), nil)
	assert.Nil(t, err)

	exists, err := VersionExists(repo, "v0.2.0")
	assert.Nil(t, err)
	assert.False(t, exists)
}
