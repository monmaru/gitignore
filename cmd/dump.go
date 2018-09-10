package cmd

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

// Dump ...
func Dump(c *cli.Context) {
	fetchGitignore(func(gi *github.Gitignore) {
		fmt.Println(gi.GetSource())
	})
}
