package check

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v6"
	"github.com/snubwoody/calver/pkg"
	"github.com/spf13/cobra"
	"strings"
)

var Cmd = &cobra.Command{
	Use:   "check",
	Short: "Check for CalVer violations",

	RunE: check,
}

var (
	ErrInvalidManifestFile = errors.New("Invalid manifest format, only package.json and Cargo.toml are supported")
)

func NewCheckCmd() *cobra.Command {
	cmd := Cmd
	cmd.Flags().StringP("manifest", "m", "", "Path to the manifest (e.g Cargo.toml)")
	cmd.Flags().StringP("repo", "r", ".", "Path to the git repository")
	return cmd
}

func check(cmd *cobra.Command, _ []string) error {
	path, err := cmd.Flags().GetString("manifest")
	if err != nil {
		return err
	}
	repoPath, err := cmd.Flags().GetString("repo")
	if err != nil {
		return err
	}

	validFormat := false
	formats := []string{"Cargo.toml", "package.json"}
	for _, format := range formats {
		if strings.HasSuffix(path, format) {
			validFormat = true
			break
		}
	}

	if !validFormat {
		return ErrInvalidManifestFile
	}

	var version string
	if strings.HasSuffix(path, "package.json") {
		p, err := pkg.ReadPackageJson(path)
		if err != nil {
			return err
		}
		version = p.Version
	}

	if strings.HasSuffix(path, "Cargo.toml") {
		p, err := pkg.ReadCargoToml(path)
		if err != nil {
			return err
		}
		version = p.Package.Version
	}

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	exists, err := pkg.VersionExists(repo, fmt.Sprintf("v%s", version))
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("version: %s already exists", version)
	}
	return nil
}
