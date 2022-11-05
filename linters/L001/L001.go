package L001

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

var LocalGomegaEventuallyAnalyzer = &analysis.Analyzer{
	Name: "L001",
	Doc:  "Checks Eventually and Consistently Gomega blocks to alway use a local Gomega variable for asserts",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	fmt.Print("L001")
	return nil, nil
}
