package base

import (
	"fmt"
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

func GetKubebuilderDefault(field *ast.Field) (string, error) {
	if field.Doc == nil {
		return "", fmt.Errorf(
			"field %v does not have kubebuilder:default comment", field)
	}
	for _, comment := range field.Doc.List {
		if strings.Contains(comment.Text, KubebuilderDefault) {
			default_ := strings.Split(comment.Text, "=")[1]
			return default_, nil
		}
	}
	return "", fmt.Errorf(
		"field %v does not have kubebuilder:default comment", field)
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
func (l *BaseLinter) ExprName(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return l.ExprName(x.X) + "." + x.Sel.Name
	case *ast.StarExpr:
		return "*" + l.ExprName(x.X)
	case *ast.CallExpr:
		return l.ExprName(x.Fun)
	case *ast.ArrayType:
		return "[]" + l.ExprName(x.Elt)
	case *ast.ParenExpr:
		return "(" + l.ExprName(x.X) + ")"
	case *ast.UnaryExpr:
		return x.Op.String() + l.ExprName(x.X)
	case *ast.CompositeLit:
		names := []string{}
		for _, expr := range x.Elts {
			names = append(names, l.ExprName(expr))
		}
		return strings.Join(names, ".")
	case *ast.FuncLit:
		return "func"
	case *ast.MapType:
		return "map[" + l.ExprName(x.Key) + "]" + l.ExprName(x.Value)
	case *ast.KeyValueExpr:
		return l.ExprName(x.Key) + ":" + l.ExprName(x.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		l.Debug(x.Pos(), "unhandled type")
		panic(fmt.Errorf("unhandled type %T", expr))
	}
}

// ExprLastName returns the last segment of the name of the expression. I..e
// "a.b.c()" result in "c". If the expression is not constructed as a
// Selector / Ident tree then it returns "".
func (l *BaseLinter) ExprLastName(expr ast.Expr) string {
	name := l.ExprName(expr)
	arr := strings.Split(name, ".")
	return arr[len(arr)-1]
}

// TypeName returns the name of the type of the field
func (l *BaseLinter) TypeName(field *ast.Field) string {
	return l.ExprName(field.Type)
}

// PointerType returns true if the type is a pointer
func PointerType(type_ ast.Expr) bool {
	_, ok := type_.(*ast.StarExpr)
	return ok
}
