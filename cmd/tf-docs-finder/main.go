//go:generate go run ../../gen/main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	tfdocsfinder "github.com/shmokmt/tf-docs-finder"
)

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
	idx, _ := fuzzyfinder.Find(tfdocsfinder.AwsResources, func(i int) string {
		return tfdocsfinder.AwsResources[i]
	})
	fmt.Println(tfdocsfinder.AwsResources[idx])
	url := fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/%s", strings.TrimPrefix(tfdocsfinder.AwsResources[idx], "aws_"))
	if err := openBrowser(url); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
