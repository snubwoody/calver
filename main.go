package main

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/go-git/go-git/v6"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:           "calver",
    Short:         "CalVer tool",
    Long:          "A tool for checking CalVer versions against git tags",
    SilenceErrors: true,
    SilenceUsage:  true,
}

var checkCmd = &cobra.Command{
    Use:   "check",
    Short: "Check for CalVer violations",
    RunE: func(cmd *cobra.Command, args []string) error {
        path, err := cmd.Flags().GetString("manifest")
        // TODO: check if manifest is missing
        if err != nil {
            return err
        }
        pkg, err := ReadPackageJson(path)
        if err != nil {
            return err
        }
        repo, err := git.PlainOpen("")
        exists, err := VersionExists(repo, fmt.Sprintf("v%s", pkg.Version))
        if err != nil {
            return err
        }
        if exists {
            return fmt.Errorf("version: %s already exists", pkg.Version)
        }
        return nil
    },
}

func Execute() {
    rootCmd.AddCommand(checkCmd)
    checkCmd.Flags().StringP("manifest", "m", "", "Path to the manifest (e.g Cargo.toml)")
    if err := rootCmd.Execute(); err != nil {
        c := color.New(color.FgRed)
        c.Fprint(os.Stderr, "Error: ")
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func main() {
    Execute()
    color.Green("No issues found")
}
