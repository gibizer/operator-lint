package main

import (
	"github.com/gibizer/operator-lint/linters/envtest/T001"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		T001.Linter,
	)
}
