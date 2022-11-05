package L001_test

import (
	"testing"

	"github.com/gibizer/envtest-linter/linters/L001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestL001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, L001.LocalGomegaEventuallyAnalyzer, "a")
}
