[![GO](https://img.shields.io/github/go-mod/go-version/obalunenko/coverbadger)](https://golang.org/doc/devel/release.html)
[![GoDoc](https://godoc.org/github.com/obalunenko/coverbadger?status.svg)](https://godoc.org/github.com/obalunenko/coverbadger)
[![Go Report Card](https://goreportcard.com/badge/github.com/obalunenko/coverbadger)](https://goreportcard.com/report/github.com/obalunenko/coverbadger)
[![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/coverbadger)](https://github.com/obalunenko/coverbadger/releases/latest)
[![License](https://img.shields.io/github/license/obalunenko/coverbadger)](/LICENSE)
[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=obalunenko_coverbadger)](https://sonarcloud.io/summary/new_code?id=obalunenko_coverbadger)
![coverbadger-tag-do-not-edit](https://img.shields.io/badge/coverage-59.68%25-brightgreen?longCache=true&style=flat)

# coverbadger
Generate coverage badge images for Markdown files using Go

## Getting Started

To install the executable (ensure your $PATH contains $GOPATH/bin):

```
go install github.com/obalunenko/coverbadger/cmd/coverbadger
```
Or download from relases [![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/coverbadger)](https://github.com/obalunenko/coverbadger/releases/latest)

## Quick Start:

<hr>
Either enter a Markdown file that does not already exist, or a Markdown file (like your README.md) that you want to update with coverage badge info.
After executing of `coverbadger` the following badge will be added

!`[coverbadger-tag-do-not-edit](<badge_url>)`

This tag will be replaced by the image for your coverage badge.

<hr>
To update a .md file badge (note: comma-separated):
Manually set the coverage value (note: do not include %):

`coverbadger -md="README.md,coverage.md" -coverage=95`

## Example:

[Example of usage](https://github.com/obalunenko/coverbadger/blob/master/scripts/update-readme-coverage.sh)


## Confused?

Try running:

```
coverbadger -h
```

## Credits

@jpoles1 - as an author of original tool [gopherbadger](https://github.com/jpoles1/gopherbadger)

