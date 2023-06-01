package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/google/go-github/v39/github"
	"github.com/ktr0731/go-fuzzyfinder"
	"golang.org/x/oauth2"
)

func normalizeResourceName(fileName string) string {
	ret := strings.TrimSuffix(fileName, ".html.markdown")
	return "aws_" + strings.TrimSuffix(ret, ".markdown")
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryContentGetOptions{}

	owner := "hashicorp"
	repo := "terraform-provider-aws"
	path := "website/docs/r"

	_, dir, _, err := client.Repositories.GetContents(ctx, owner, repo, path, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resources []string
	for _, f := range dir {
		name := normalizeResourceName(f.GetName())
		resources = append(resources, name)
	}
	idx, _ := fuzzyfinder.Find(resources, func(i int) string {
		return resources[i]
	})
	url := fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/%s", strings.TrimPrefix(resources[idx], "aws_"))
	if err := openBrowser(url); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
