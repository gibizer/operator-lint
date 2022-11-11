package tools

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Report a lint rule violation
func Report(pass *analysis.Pass, pos token.Pos, analyzerName string, format string, a ...any) {
	pass.Reportf(pos, analyzerName+": "+format, a...)
}

// HasDocComment returns true if the field has a doc comment containting the
// substring
func HasDocComment(field *ast.Field, substring string) bool {
	if field.Doc == nil {
		return false
	}
	for _, comment := range field.Doc.List {
		if strings.Contains(comment.Text, substring) {
			return true
		}
	}
	return false
}
