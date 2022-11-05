package L001

import (
	"golang.org/x/tools/go/analysis"
)

var LocalGomegaEventuallyAnalyzer = &analysis.Analyzer{
	Name: "Local Gomega analyzer",
	Doc:  "Checks Eventually and Consistently Gomega blocks to alway use a local Gomega variable for asserts",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
