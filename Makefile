# HELP sourced from https://gist.github.com/prwhite/8168133

# Add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_FUNC = \
	%help; \
	while(<>) { \
		if(/^([a-z0-9_-]+):.*\#\#(?:@(\w+))?\s(.*)$$/) { \
			push(@{$$help{$$2}}, [$$1, $$3]); \
		} \
	}; \
	print "usage: make [target]\n\n"; \
	for ( sort keys %help ) { \
		print "$$_:\n"; \
		printf("  %-30s %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
		print "\n"; \
	}

help: ##@help show this help
	@perl -e '$(HELP_FUNC)' $(MAKEFILE_LIST)

GO111MODULE=on
GOSUMDB=sum.golang.org
export GO111MODULE
export GOSUMDB

COVER_PACKAGE=$(shell go list ./... | grep -v -e "/testfactory" -e "/mocks" | paste -s -d ',' -)
MIN_COVERAGE=0

APP_EXECUTABLE="./out/portal"
GIT_VERSION?=$(shell git describe --always --tags --long --dirty --abbrev=40)
GO_LDFLAGS="-X github.com/techx/portal/version.Version=$(GIT_VERSION)"

TOOLS_MOD_DIR = ./tools
TOOLS_DIR = $(abspath ./.tools)

$(TOOLS_DIR)/golangci-lint: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(TOOLS_DIR)/gotest: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/gotest github.com/rakyll/gotest

$(TOOLS_DIR)/mockery: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/mockery github.com/vektra/mockery/v2

$(TOOLS_DIR)/gocover-cobertura: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/gocover-cobertura github.com/boumenot/gocover-cobertura

$(TOOLS_DIR)/swagger: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/swagger github.com/go-swagger/go-swagger/cmd/swagger

$(TOOLS_DIR)/gofumpt: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/gofumpt mvdan.cc/gofumpt

$(TOOLS_DIR)/gci: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/gci github.com/daixiang0/gci

$(TOOLS_DIR)/goawk: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/goawk github.com/benhoyt/goawk

$(TOOLS_DIR)/go-junit-report: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/go-junit-report github.com/jstemmer/go-junit-report/v2


# DEVELOPMENT	###########################################################################################

run-local: ##@development runs a local server directly from source
	go run cmd/portal/*.go start --config-file test.application.yml

sample-config: ##@development generates sample.application.yml
	go run cmd/portal/*.go generate-config --config-file sample.application.yml

lint: $(TOOLS_DIR)/golangci-lint ##@development runs linting using golangci-lint
	$(TOOLS_DIR)/golangci-lint run --config=.golangci.yml

generate-swagger: $(TOOLS_DIR)/swagger ##@development generates swagger docs
	$(TOOLS_DIR)/swagger generate spec -m -o swagger/v1/swagger.yaml ./cmd/portal -x github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options

generate-mocks: $(TOOLS_DIR)/mockery delete-mocks ##@development generates mock for all existing interfaces
	$(TOOLS_DIR)/mockery

delete-mocks:
	rm $(shell find . | grep mocks/ | grep -v -e with_recorder)

fumpt: $(TOOLS_DIR)/gofumpt
	$(TOOLS_DIR)/gofumpt -l -w -extra .

imports: $(TOOLS_DIR)/gci
	$(TOOLS_DIR)/gci write ./ -s standard -s default --skip-generated

fmt: fumpt imports ##@development formats code

lsif: $(TOOLS_DIR)/lsif-go
	@mkdir -p lsif
	@$(TOOLS_DIR)/lsif-go --verbose --no-animation --output="lsif/dump.lsif"

# TESTS 		###########################################################################################

test: $(TOOLS_DIR)/gotest ##@tests runs tests
	go mod tidy
	$(TOOLS_DIR)/gotest -race -v ./...

test-coverage: $(TOOLS_DIR)/gocover-cobertura $(TOOLS_DIR)/goawk ##@tests generates coverage report
	mkdir -p coverage
	go test -mod=readonly -race  -cover -coverprofile=coverage/coverage.out -covermode atomic -coverpkg=$(COVER_PACKAGE) ./...
	@go tool cover -func=coverage/coverage.out
	@go tool cover -func=coverage/coverage.out | .tools/goawk  '/total:.*statements/ {if ($$3 < $(MIN_COVERAGE)) {print "ERR: coverage is lower than $(MIN_COVERAGE)"; exit 1}}'

test-coverage-html: test-coverage ##@tests generates coverage report as html
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	open coverage/coverage.html

.PHONY: coverage lint test

# CI			###########################################################################################

ci: build lint test-coverage ##@ci runs what CI runs

ci-lint: $(TOOLS_DIR)/golangci-lint ##@ci runs linting using golangci-lint for ci
	$(TOOLS_DIR)/golangci-lint run --config=.golangci.yml --out-format code-climate > code-quality-report.json \
		|| (cat code-quality-report.json | jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'; exit 1)

ci-test: $(TOOLS_DIR)/gocover-cobertura $(TOOLS_DIR)/go-junit-report $(TOOLS_DIR)/goawk ##@tests generates coverage report
	mkdir -p coverage
	mkdir -p testreport
	go test -mod=readonly -v -race -cover -coverprofile=coverage/coverage.out -covermode atomic -coverpkg=$(COVER_PACKAGE) ./... 2>&1 \
		| tee testreport/test.out
	@go tool cover -func=coverage/coverage.out
	$(TOOLS_DIR)/gocover-cobertura < coverage/coverage.out > coverage/cobertura-coverage.xml
	$(TOOLS_DIR)/go-junit-report -set-exit-code -in testreport/test.out -out testreport/junit.xml
	@go tool cover -func=coverage/coverage.out | .tools/goawk  '/total:.*statements/ {if ($$3 < $(MIN_COVERAGE)) {print "ERR: coverage is lower than $(MIN_COVERAGE)"; exit 1}}'

clean:
	go clean cmd/portal/*.go
	rm -rf out/
	go mod tidy

build: clean
	mkdir -p out
	go build -o $(APP_EXECUTABLE) -ldflags $(GO_LDFLAGS) cmd/portal/*.go
