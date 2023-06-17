package base_test

import (
	"fmt"
	"go/ast"
	"testing"

	"github.com/gibizer/operator-lint/pkg/base"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

type Linter struct {
	*base.BaseLinter
}

func NewAnalyzer() *analysis.Analyzer {
	l := &Linter{}
	l.BaseLinter = base.NewBaseLinter("TestLint", "Test doc", l)
	return l.Analyzer
}

func (l *Linter) LintFile(file *ast.File) error {
	return fmt.Errorf("Should have been skipped")
}

func TestSkipFlag(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := NewAnalyzer()
	err := analyzer.Flags.Parse([]string{"-skip"})
	if err != nil {
		t.Fail()
	}

	analysistest.Run(t, testdata, analyzer, "a")
}
