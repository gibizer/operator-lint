package C001_test

import (
	"testing"

	"github.com/gibizer/operator-lint/linters/crd/C001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestC001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, C001.NewAnalyzer(), "a")
}
