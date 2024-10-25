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
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "type",
			Value: "resources",
			Usage: "resources or data-sources. default is resources.",
		},
	},
	Action: func(cCtx *cli.Context) error {
		tfType := cCtx.String("type")
		var list []string
		switch tfType {
		case "resources":
			list = tfdocsfinder.GoogleResources
		case "data-sources":
			list = tfdocsfinder.GoogleDataSources
		default:
			return cli.Exit(fmt.Errorf("invalid type: %s", tfType), 1)
		}
		idx, err := fuzzyfinder.Find(list, func(i int) string {
			return list[i]
		})
		if err != nil {
			return nil
		}
		url := fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/google/latest/docs/%s/%s", tfType, strings.TrimPrefix(list[idx], "google_"))
		if err := openBrowser(url); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
