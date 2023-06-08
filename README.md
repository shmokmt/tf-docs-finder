# tf-docs-finder

This utility tool is designed to immediately open the documentation of a Terraform Provider from the command line.

## Installation

```
go install github.com/shmokmt/tf-docs-finder/cmd/tf-docs-finder@latest
```

## Usage

### AWS Provider

List resources:

```
tf-docs-finder aws
```

List data-sources

```
tf-docs-finder aws --type data-sources
```

Other providers have the same interface as well.

## Terraform Provider

Available following providers:

- [hashicorp/google](https://registry.terraform.io/providers/hashicorp/google/latest/docs)
- [hashicorp/aws](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
- [integrations/github](https://registry.terraform.io/providers/integrations/github/latest/docs)

At present, it only supports latest version.
