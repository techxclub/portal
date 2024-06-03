//go:build tools
// +build tools

package tools

import (
	_ "github.com/benhoyt/goawk"
	_ "github.com/boumenot/gocover-cobertura"
	_ "github.com/daixiang0/gci"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/jstemmer/go-junit-report/v2"
	_ "github.com/rakyll/gotest"
	_ "github.com/sourcegraph/lsif-go/cmd/lsif-go"
	_ "github.com/vektra/mockery/v2"
	_ "mvdan.cc/gofumpt"
)
