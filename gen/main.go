package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var camelCaseMap = map[string]string{
	"aws":    "Aws",
	"github": "GitHub",
	"google": "Google",
}

type generator struct {
	gh *github.Client
}

func newGenerator(ctx context.Context) *generator {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &generator{
		gh: github.NewClient(tc),
	}
}

func (g *generator) run(ctx context.Context, p *provider) error {
	commits, _, err := g.gh.Repositories.ListCommits(ctx, p.repo.owner, p.repo.name, &github.CommitsListOptions{})
	if err != nil {
		return err
	}
	commitSHA := *commits[0].SHA

	tree, _, err := g.gh.Git.GetTree(ctx, p.repo.owner, p.repo.name, commitSHA, true)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.WriteString("// Code generated by gen/main.go; DO NOT EDIT.")
	buf.WriteString("\n\npackage tfdocsfinder")
	buf.WriteString(fmt.Sprintf("\nvar %sResources = []string{", camelCaseMap[p.name]))

	var resourceNames []string
	for _, entry := range tree.Entries {
		if *entry.Type == "blob" && strings.HasPrefix(*entry.Path, p.repo.resourcesPath) {
			path := strings.TrimPrefix(*entry.Path, p.repo.resourcesPath)
			path = strings.TrimSuffix(path, ".html.markdown")
			name := p.name + "_" + strings.TrimSuffix(path, ".markdown")
			resourceNames = append(resourceNames, name)
			buf.WriteString("\n\t\"" + name + "\",")
		}
	}
	buf.WriteString("\n}")
	fmt.Println(len(resourceNames))

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("../../%s_resources.go", p.name)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		return err
	}

	buf.Reset()
	buf.WriteString("// Code generated by gen/main.go; DO NOT EDIT.")
	buf.WriteString("\n\npackage tfdocsfinder")
	buf.WriteString(fmt.Sprintf("\nvar %sDataSources = []string{", camelCaseMap[p.name]))

	for _, entry := range tree.Entries {
		if *entry.Type == "blob" && strings.HasPrefix(*entry.Path, p.repo.dataSourcesPath) {
			path := strings.TrimPrefix(*entry.Path, p.repo.dataSourcesPath)
			path = strings.TrimSuffix(path, ".html.markdown")
			name := p.name + "_" + strings.TrimSuffix(path, ".markdown")
			buf.WriteString("\n\t\"" + name + "\",")
		}
	}
	buf.WriteString("\n}")

	formatted, err = format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	fileName = fmt.Sprintf("../../%s_data_sources.go", p.name)
	f, err = os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		return err
	}
	return nil
}

type provider struct {
	name string
	repo *repository
}

type repository struct {
	name            string
	owner           string
	resourcesPath   string
	dataSourcesPath string
}

var providers = []*provider{
	{
		name: "aws",
		repo: &repository{
			name:            "terraform-provider-aws",
			owner:           "hashicorp",
			resourcesPath:   "website/docs/r/",
			dataSourcesPath: "website/docs/d/",
		},
	},
	{
		name: "github",
		repo: &repository{
			name:            "terraform-provider-github",
			owner:           "integrations",
			resourcesPath:   "website/docs/r/",
			dataSourcesPath: "website/docs/d/",
		},
	},
	{
		name: "google",
		repo: &repository{
			name:            "terraform-provider-google",
			owner:           "hashicorp",
			resourcesPath:   "website/docs/r/",
			dataSourcesPath: "website/docs/d/",
		},
	},
}

func main() {
	generator := newGenerator(context.Background())
	for _, provider := range providers {
		err := generator.run(context.Background(), provider)
		if err != nil {
			log.Fatal(err)
		}
	}
}
