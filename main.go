package main

import (
	"github.com/gibizer/envtest-linter/linters/L001"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		L001.LocalGomegaEventuallyAnalyzer,
	)
}
