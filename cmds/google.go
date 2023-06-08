package cmds

import (
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	tfdocsfinder "github.com/shmokmt/tf-docs-finder"
	"github.com/urfave/cli/v2"
)

var GoogleCommand = &cli.Command{
	Name:        "google",
	Description: "Open hashicorp/google's docs",
	Action: func(cCtx *cli.Context) error {
		idx, err := fuzzyfinder.Find(tfdocsfinder.GoogleResources, func(i int) string {
			return tfdocsfinder.GoogleResources[i]
		})
		if err != nil {
			return nil
		}
		fmt.Println(tfdocsfinder.GoogleResources[idx])
		url := fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/%s", strings.TrimPrefix(tfdocsfinder.GoogleResources[idx], "google_"))
		if err := openBrowser(url); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
