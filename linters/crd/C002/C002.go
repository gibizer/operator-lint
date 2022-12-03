package C002

import (
	"go/ast"
	"strings"

	b "github.com/gibizer/operator-lint/pkg/base"
	"golang.org/x/tools/go/analysis"
)

const (
	Name = "C002"
	Doc  = "detects incompatible `Required` kubebuilder marker and `omitemty` golang tag"
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
		switch x := node.(type) {
		case *ast.Field:
			if x.Tag == nil {
				return true
			}
			required := b.HasDocComment(x, "+kubebuilder:validation:Required")
			if required && strings.Contains(x.Tag.Value, ",omitempty") {
				l.Report(
					x.Pos(),
					"Field '%s' has both a 'Required' kubebuilder marker and an 'omitempty' tag. "+
						"Only one of them should be used", x.Names[0].Name)
			}
		}
		return true
	})
	return nil
}
