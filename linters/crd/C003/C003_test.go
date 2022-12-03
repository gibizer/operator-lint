package C003_test

import (
	"testing"

	"github.com/gibizer/operator-lint/linters/crd/C003"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestC003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, C003.NewAnalyzer(), "a")
}
