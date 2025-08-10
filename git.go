package calver_check

import (
	"errors"
	"github.com/go-git/go-git/v6"
)

// Checks if a git tag for the version exists.
func VersionExists(repo *git.Repository, version string) (bool, error) {
	_, err := repo.Tag(version)
	if errors.Is(err, git.ErrTagNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}
