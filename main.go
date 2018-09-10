package main

import (
	"fmt"
	"os"

	"github.com/monmaru/gitignore/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitignore"
	app.Version = "1.0"
	app.Author = "monmaru"
	app.Description = "generate a .gitignore from github/gitignore on GitHub"
	app.Commands = commands
	app.CommandNotFound = usage
	app.Run(os.Args)
}

var commands = []cli.Command{
	{
		Name:   "list",
		Usage:  "List a collection of .gitignore templates",
		Action: cmd.List,
	},
	{
		Name:   "dump",
		Usage:  "Dump a .gitignore to console",
		Action: cmd.Dump,
	},
	{
		Name:   "new",
		Usage:  "Create a .gitignore to current directory",
		Action: cmd.New,
	},
}

func usage(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
