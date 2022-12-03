package C003

import (
	"go/ast"
	"strings"

	b "github.com/gibizer/operator-lint/pkg/base"
	"golang.org/x/tools/go/analysis"
)

const (
	Name = "C003"
	Doc  = "detects incompatible defaulting via `Optional` kubebuilder marker and `omitemty` golang tag"
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
			optional := b.HasDocComment(x, "+kubebuilder:validation:Optional")
			default_ := b.HasDocComment(x, "+kubebuilder:default")
			if optional && default_ && strings.Contains(x.Tag.Value, ",omitempty") {
				l.Report(
					x.Pos(),
					"Field '%s' has both a 'Optional' kubebuilder marker with a default value "+
						"and an 'omitempty' tag. Either remove the default value or remove 'omitempty'",
					x.Names[0].Name)
			}
		}
		return true
	})
	return nil
}
