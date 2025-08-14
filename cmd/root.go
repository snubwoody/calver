package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/snubwoody/calver/cmd/check"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:           "calver",
	Short:         "CalVer tool",
	Long:          "A tool for checking CalVer versions against git tags",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	RootCmd.AddCommand(check.NewCheckCmd())
	if err := RootCmd.Execute(); err != nil {
		c := color.New(color.FgRed)
		c.Fprint(os.Stderr, "Error: ")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
