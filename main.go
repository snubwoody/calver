package main

import (
	"github.com/fatih/color"
	"github.com/snubwoody/calver/cmd"
)

func main() {
	cmd.Execute()
	color.Green("No issues found")
}
