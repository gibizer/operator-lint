package T001_test

import (
	"testing"

	"github.com/gibizer/operator-lint/linters/envtest/T001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestL001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, T001.Linter, "a")
}
