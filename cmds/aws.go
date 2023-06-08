package cmds

import (
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	tfdocsfinder "github.com/shmokmt/tf-docs-finder"
	"github.com/urfave/cli/v2"
)

var AwsCommand = &cli.Command{
	Name:        "aws",
	Description: "Open hashicorp/aws's docs",
	Action: func(cCtx *cli.Context) error {
		idx, err := fuzzyfinder.Find(tfdocsfinder.AwsResources, func(i int) string {
			return tfdocsfinder.AwsResources[i]
		})
		if err != nil {
			return nil
		}
		fmt.Println(tfdocsfinder.AwsResources[idx])
		url := fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/%s", strings.TrimPrefix(tfdocsfinder.AwsResources[idx], "aws_"))
		if err := openBrowser(url); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
