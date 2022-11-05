module github.com/gibizer/envtest-linter

go 1.18

// All versions here need to be the same as in golangci-lint/mod.go if present
require golang.org/x/tools v0.0.0-20191010075000-0337d82405ff

replace golang.org/x/tools => github.com/golangci/tools v0.0.0-20190915081210-016959b1a823
