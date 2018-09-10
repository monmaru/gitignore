package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

// List ...
func List(c *cli.Context) {
	ctx := context.Background()
	client := github.NewClient(nil)

	ignores, _, err := client.Gitignores.List(ctx)
	check(err)

	fmt.Println("--------------------------------------")
	fmt.Println("A collection of .gitignore templates")
	fmt.Println("--------------------------------------")
	for _, ignore := range ignores {
		fmt.Println("  " + ignore)
	}
	fmt.Println()
}
