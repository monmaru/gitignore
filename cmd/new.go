package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

// New ...
func New(c *cli.Context) {
	fetchGitignore(func(gi *github.Gitignore) {
		f := ".gitignore"
		if _, err := os.Stat(f); err == nil {
			fmt.Println("--------------------------------------")
			fmt.Printf("%s exists\n", f)
			fmt.Println("--------------------------------------")
			return
		}

		check(ioutil.WriteFile(f, []byte(gi.GetSource()), 0666))
	})
}
