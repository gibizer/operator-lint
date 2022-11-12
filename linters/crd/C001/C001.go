package C001

import (
	"go/ast"
	"strings"

	b "github.com/gibizer/operator-lint/pkg/base"
	"golang.org/x/tools/go/analysis"
)

const (
	Name = "C001"
	Doc  = "Checks incompatible kubebuilder markers and golang tags on CRD fields"
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
			required := b.HasDocComment(x, "+kubebuilder:validation:Required")
			optional := b.HasDocComment(x, "+kubebuilder:validation:Optional")
			if required && optional {
				l.Report(
					x.Pos(),
					"Field '%s' has both 'Optional' and 'Required' kubebuilder markers. "+
						"Only one of them should be used", x.Names[0].Name)
			}
			if x.Tag == nil {
				return true
			}
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
