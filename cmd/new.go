package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

// New creates a .gitignore to current directory
func New(c *cli.Context) {
	fetchGitignore(func(gi *github.Gitignore) {
		f := ".gitignore"
		if _, err := os.Stat(f); err == nil {
			fmt.Println("--------------------------------------")
			fmt.Printf("%s already exists\n", f)
			fmt.Println("--------------------------------------")
			return
		}

		check(ioutil.WriteFile(f, []byte(gi.GetSource()), 0666))
		fmt.Println("--------------------------------------")
		fmt.Printf("%s was created\n", f)
		fmt.Println("--------------------------------------")
	})
}
