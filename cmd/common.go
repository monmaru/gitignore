package cmd

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

func fetchGitignore(fn func(*github.Gitignore)) {
	ctx := context.Background()
	client := github.NewClient(nil)

	ignores, _, err := client.Gitignores.List(ctx)
	check(err)

	name := ignorePrompt(ignores)
	gi, _, err := client.Gitignores.Get(ctx, name)
	check(err)

	fn(gi)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
