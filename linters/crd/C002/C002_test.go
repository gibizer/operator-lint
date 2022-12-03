package C002_test

import (
	"testing"

	"github.com/gibizer/operator-lint/linters/crd/C002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestC002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, C002.NewAnalyzer(), "a")
}
