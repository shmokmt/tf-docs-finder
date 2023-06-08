//go:generate go run ../../gen/main.go
package main

import (
	"log"
	"os"

	"github.com/shmokmt/tf-docs-finder/cmds"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tf-docs-finder",
		Usage: "tf-docs-finder",
		Commands: []*cli.Command{
			cmds.AwsCommand,
			cmds.GitHubCommand,
			cmds.GoogleCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
