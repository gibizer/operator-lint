package FIXME

import (
	"go/ast"

	b "github.com/gibizer/operator-lint/pkg/base"
	"golang.org/x/tools/go/analysis"
)

const (
	Name = "FIXME"
	Doc  = "<descibe the check>"
)

type Linter struct {
	*b.BaseLinter
}

func NewAnalyzer() *analysis.Analyzer {
	l := &Linter{}
	l.BaseLinter = b.NewBaseLinter(Name, Doc, l)
	return l.Analyzer
}

func (l *Linter) LintFile(file *ast.File) error {
	ast.Inspect(file, func(node ast.Node) bool {
		return true
	})
	return nil
}
