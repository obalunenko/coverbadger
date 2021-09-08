![coverbadger-tag-do-not-edit](https://img.shields.io/badge/coverage-55.56%25-brightgreen?longCache=true&style=flat)

# coverbadger
Generate coverage badge images for Markdown files using Go

## Getting Started

To install the executable (ensure your $PATH contains $GOPATH/bin):

```
go install github.com/obalunenko/coverbadger
```

## Quick Start:

<hr>
Either enter a Markdown file that does not already exist, or a Markdown file (like your README.md) that contains the following tag somewhere in the contents:

!`[coverbadger-tag-do-not-edit]()`

This tag will be replaced by the image for your coverage badge.

<hr>
To update a .md file badge (note: comma-separated):
Manually set the coverage value (note: do not include %):

`coverbadger -md="README.md,coverage.md" -coverage=95`


## Confused?

Try running:

```
coverbadger -h
```

