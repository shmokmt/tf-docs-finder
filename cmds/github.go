package cmds

import (
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	tfdocsfinder "github.com/shmokmt/tf-docs-finder"
	"github.com/urfave/cli/v2"
)

var GitHubCommand = &cli.Command{
	Name:        "github",
	Description: "Open integrations/github's docs",
	Action: func(cCtx *cli.Context) error {
		idx, err := fuzzyfinder.Find(tfdocsfinder.GitHubResources, func(i int) string {
			return tfdocsfinder.GitHubResources[i]
		})
		if err != nil {
			return nil
		}
		fmt.Println(tfdocsfinder.GitHubResources[idx])
		url := fmt.Sprintf("https://registry.terraform.io/providers/integrations/github/latest/docs/resources/%s", strings.TrimPrefix(tfdocsfinder.GitHubResources[idx], "github_"))
		if err := openBrowser(url); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
