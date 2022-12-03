package FIXME_test

import (
	"testing"

	"github.com/gibizer/operator-lint/lint-template/FIXME"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestFIXME(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, FIXME.NewAnalyzer(), "a")
}
