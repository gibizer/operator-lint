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

var typeNameToDefaultValue = map[string]string{
	"int32":  "0",
	"int":    "0",
	"string": "",
	"bool":   "false",
}

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
			default_ := b.HasDocComment(x, b.KubebuilderDefault)

			if !(optional && default_) {
				return true
			}

			omitempty := strings.Contains(x.Tag.Value, ",omitempty")
			if !omitempty {
				return true
			}

			if b.PointerType(x.Type) {
				return true
			}

			typeName := l.TypeName(x)
			golangDefault, ok := typeNameToDefaultValue[typeName]
			if !ok {
				l.Report(
					x.Pos(),
					"Field '%s' has both a 'Optional' kubebuilder marker with a default value "+
						"and an 'omitempty' tag. Either remove the default value or remove 'omitempty' or change the "+
						"field to a pointer type",
					x.Names[0].Name)

				return true
			}

			defaultValue, err := b.GetKubebuilderDefault(x)
			if err != nil {
				panic(err)
			}

			if defaultValue != golangDefault {
				l.Report(
					x.Pos(),
					"Field '%s' has both a 'Optional' kubebuilder marker with a default value "+
						"and an 'omitempty' tag. Either remove the default value or remove 'omitempty' or change the "+
						"field to a pointer type",
					x.Names[0].Name)
				return true
			}

		}
		return true
	})
	return nil
}
