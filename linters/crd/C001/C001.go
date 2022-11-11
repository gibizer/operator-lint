package C001

import (
	"go/ast"
	"strings"

	tools "github.com/gibizer/operator-lint/pkg"
	"golang.org/x/tools/go/analysis"
)

const (
	Name = "C001"
	Doc  = "cchecks incompatible kubebuilder markers and golang tags on CRD fields"
)

var Linter = &analysis.Analyzer{
	Name: Name,
	Doc:  Doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.Field:
				required := tools.HasDocComment(x, "+kubebuilder:validation:Required")
				optional := tools.HasDocComment(x, "+kubebuilder:validation:Optional")
				if required && optional {
					tools.Report(
						pass, x.Pos(), Name,
						"Field '%s' has both 'Optional' and 'Required' kubebuilder markers. "+
							"Only one of them should be used", x.Names[0].Name)
				}
				if x.Tag == nil {
					return true
				}
				if required && strings.Contains(x.Tag.Value, ",omitempty") {
					tools.Report(
						pass, x.Pos(), Name,
						"Field '%s' has both a 'Required' kubebuilder marker and an 'omitempty' tag. "+
							"Only one of them should be used", x.Names[0].Name)
				}
			}
			return true
		})
	}
	return nil, nil
}
