.DEFAULT_GOAL := help

BENCH_TARGET := .

help: ## Displays this help message.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: vet ## Runs the tests and vetting
	go test -v -race -count=1 ./...

vet: ## Run vetting checks
	go vet ./...

semgrep: ## Run semgrep
	semgrep --config auto

lint: ## Run golangci-lint
	golangci-lint run ./...

report: ## Run goreportcard
	goreportcard-cli

test-all: semgrep lint test report ## Run the tests, vetting, and golangci-lint, and semgrep

ci: vet ## Run the tests and vetting checks (specific for CI)
	go test -cover -race -count=1 ./...

cover: FILE := coverage.txt
cover: ## Run coverage
	go test -race -coverprofile=$(FILE) -covermode=atomic $(TARGET)
	go tool cover -func=$(FILE)

cover-all: ## Run coverage but recursive
cover-all: TARGET := ./...
cover-all: cover

bench: ## Run benchmarking
	go test -run "^$$" -bench "$(BENCH_TARGET)" -cpu=4 $(ARGS)

bench-initialize: ## Run benchmarking for initializing
bench-initialize: BENCH_TARGET := ^BenchmarkInitialize
bench-initialize: bench

bench-analyze: ## Run benchmarking for analyzing
bench-analyze: BENCH_TARGET := ^BenchmarkAnalyze
bench-analyze: bench

bench-analyze-dev: ## Run benchmarking for analyzing (uncached)
bench-analyze-dev: BENCH_TARGET := ^BenchmarkAnalyze.+WithDevelopment$
bench-analyze-dev: bench

licensing: ## Run licensing Go files
	go-license --config .github/license.yaml $(ARGS) *.go **/*.go

license-verify: ## Verifying files license
license-verify: ARGS := --verify
license-verify: licensing

pprof: ## Run pprof
pprof: ARGS := -cpuprofile=cpu.out -memprofile=mem.out
pprof: bench