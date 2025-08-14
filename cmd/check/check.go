package check

import (
	"fmt"
	"github.com/go-git/go-git/v6"
	"github.com/snubwoody/calver/pkg"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for CalVer violations",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("manifest")
		// TODO: check if manifest is missing
		if err != nil {
			return err
		}
		p, err := pkg.ReadPackageJson(path)
		if err != nil {
			return err
		}
		repo, err := git.PlainOpen("")
		exists, err := pkg.VersionExists(repo, fmt.Sprintf("v%s", p.Version))
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("version: %s already exists", p.Version)
		}
		return nil
	},
}
