package T001

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	b "github.com/gibizer/operator-lint/pkg/base"
)

const (
	Name = "L001"
	Doc  = "Checks Eventually and Consistently Gomega blocks to alway use a local Gomega variable for asserts"
)

type Linter struct {
	*b.BaseLinter
}

func NewAnalyzer() *analysis.Analyzer {
	l := &Linter{}
	l.BaseLinter = b.NewBaseLinter(Name, Doc, l)
	return l.Analyzer
}

const (
	GomegaPkg  = "\"github.com/onsi/gomega\""
	GomegaType = "Gomega"
)

func (l *Linter) LintFile(file *ast.File) error {

	imp := b.ImportSpec(file, GomegaPkg)
	if imp == nil {
		return nil
	}

	var gomegaTypeName string
	if imp.Name == nil {
		gomegaTypeName = "gomega." + GomegaType
	} else if imp.Name.Name == "." {
		gomegaTypeName = GomegaType
	} else {
		gomegaTypeName = imp.Name.Name + "." + GomegaType
	}

	ast.Inspect(file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.CallExpr:
			callName := l.ExprLastName(x.Fun)
			if callName != "Eventually" && callName != "Consistently" {
				return true
			}
			anonF := firstArgAsFuncLit(x)
			if anonF == nil {
				// If the first arg is not a function then this is
				// a call with value that we can ignore here
				// https://onsi.github.io/gomega/#category-1-making-codeeventuallycode-assertions-on-values
				// Or a non anonymous function is passed. In that case
				// we would need to look up that function to analyse it. But
				// that is a TODO.
				return true
			}

			gomegaArgName := l.gomegaArgName(anonF, gomegaTypeName)

			l.checkExpectCalls(anonF.Body, gomegaArgName, callName)
			// need to go deeper as there can be nested Eventually blocks
			return true
		}
		return true
	})
	return nil
}

func firstArgAsFuncLit(f *ast.CallExpr) *ast.FuncLit {
	if len(f.Args) == 0 {
		return nil
	}
	flit, ok := f.Args[0].(*ast.FuncLit)
	if !ok {
		return nil
	}
	return flit
}

// gomegaArgName returns the name of the first Gomega argument of the function
// returns "" if the function has no such argument
func (l *Linter) gomegaArgName(f *ast.FuncLit, gomegaTypeName string) string {
	for _, param := range f.Type.Params.List {
		if l.ExprName(param.Type) == gomegaTypeName {
			return param.Names[0].Name
		}
	}
	return ""
}

func (l *Linter) checkExpectCalls(body *ast.BlockStmt, gomegaArgName string, blockName string) {
	ast.Inspect(body, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.CallExpr:
			callName := l.ExprName(x.Fun)
			if strings.HasSuffix(callName, "Eventually") || strings.HasSuffix(callName, "Consistently") {
				// stop checking deeper based on the current block as there is
				// a new nested Eventually / Consistently block. This
				// nested block will be handled by the caller.
				return false
			}

			if !strings.HasSuffix(callName, ".Expect") && callName != "Expect" {
				return true
			}

			if gomegaArgName == "" {
				l.Report(
					x.Fun.Pos(),
					"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Change the "+
						"declaration of the function passed to the '%s' block to take a parameter with type "+
						"'Gomega' and use that to call 'Expect'", blockName, blockName,
				)
				return false
			}
			if !strings.HasPrefix(callName, gomegaArgName+".") {
				l.Report(
					x.Fun.Pos(),
					"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Use '%s.Expect'",
					blockName, gomegaArgName,
				)
				return false
			}
		}
		return true
	})
}
