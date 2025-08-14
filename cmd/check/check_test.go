package check

import (
	"github.com/BurntSushi/toml"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/snubwoody/calver/pkg"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEmptyStringManifest(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().StringP("manifest", "m", "", "")
	cmd.Flags().StringP("repo", "r", ".", "")
	err := check(&cmd, []string{})
	assert.ErrorIs(t, err, ErrInvalidManifestFile)
}

func TestSupportedManifest(t *testing.T) {
	dir, err := os.MkdirTemp("", "test-output-")
	assert.Nil(t, err)
	defer os.RemoveAll(dir)

	m := pkg.CargoToml{
		Package: pkg.CargoPackage{
			Version: "24.0.0",
		},
	}

	b, err := toml.Marshal(m)
	assert.Nil(t, err)
	path := filepath.Join(dir, "Cargo.toml")
	err = os.WriteFile(path, b, 0777)
	assert.Nil(t, err)

	repo, err := git.PlainInit(dir, false)
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

	cmd := cobra.Command{}
	cmd.Flags().StringP("manifest", "m", path, "")
	cmd.Flags().StringP("repo", "r", dir, "")
	err = check(&cmd, []string{})
	assert.Nil(t, err)
}
