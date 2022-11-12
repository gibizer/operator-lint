package base

import (
	"go/ast"
	"strings"
)

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

func ImportSpec(file *ast.File, pkg string) *ast.ImportSpec {
	for _, imp := range file.Imports {
		if imp.Path.Value == pkg {
			return imp
		}
	}
	return nil
}

// ExprName returns the expression as a name. I.e. "a.b.c()"" will result in
// "a.b.c" and x() will result in "x". If expression is not constructed as
// a Selector / Ident tree then it returns "".
func ExprName(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return ExprName(x.X) + "." + x.Sel.Name
	}
	return ""
}

// ExprLastName returns the last segment of the name of the expression. I..e
// "a.b.c()" result in "c". If the expression is not constructed as a
// Selector / Ident tree then it returns "".
func ExprLastName(expr ast.Expr) string {
	name := ExprName(expr)
	arr := strings.Split(name, ".")
	return arr[len(arr)-1]
}
