package main

import (
	"github.com/gibizer/operator-lint/linters/crd/C001"
	"github.com/gibizer/operator-lint/linters/crd/C002"
	"github.com/gibizer/operator-lint/linters/crd/C003"
	"github.com/gibizer/operator-lint/linters/envtest/T001"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		T001.NewAnalyzer(),
		C001.NewAnalyzer(),
		C002.NewAnalyzer(),
		C003.NewAnalyzer(),
	)
}
