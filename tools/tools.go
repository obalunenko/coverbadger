//go:build tools

package tools

//go:generate go clean

//go:generate go install -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint

//go:generate go install -mod=vendor github.com/vasi-stripe/gogroup/cmd/gogroup

//go:generate go install -mod=vendor gotest.tools/gotestsum

//go:generate go install -mod=vendor github.com/axw/gocov/gocov

//go:generate go install -mod=vendor github.com/matm/gocov-html

//go:generate go install -mod=vendor golang.org/x/tools/cmd/cover

//go:generate go install -mod=vendor golang.org/x/tools/cmd/goimports

//go:generate go install -mod=vendor github.com/goreleaser/goreleaser

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/matm/gocov-html"
	_ "github.com/vasi-stripe/gogroup/cmd/gogroup"
	_ "golang.org/x/tools/cmd/cover"
	_ "golang.org/x/tools/cmd/goimports"
	_ "gotest.tools/gotestsum"
)
