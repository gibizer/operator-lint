package T001

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"

	tools "github.com/gibizer/operator-lint/pkg"
)

const (
	Name = "L001"
	Doc  = "Checks Eventually and Consistently Gomega blocks to alway use a local Gomega variable for asserts"
)

var Linter = &analysis.Analyzer{
	Name: Name,
	Doc:  Doc,
	Run:  run,
}

const (
	GomegaImport = "\"github.com/onsi/gomega\""
)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {

		var gomegaImportName string
		var imported = false
		for _, imp := range file.Imports {
			if imp.Path.Value == GomegaImport {
				imported = true
				//				filename := pass.Fset.Position(imp.Pos()).Filename
				//				log.Printf("L001 on file %s", filename)
				if imp.Name == nil {
					gomegaImportName = "gomega"
				} else if imp.Name.Name == "." {
					gomegaImportName = ""
				} else {
					gomegaImportName = imp.Name.Name
				}
			}
		}

		if !imported {
			continue
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.CallExpr:
				var callName string
				switch call := x.Fun.(type) {
				case *ast.Ident:
					if call.Name != "Eventually" && call.Name != "Consistently" {
						return true
					}
					callName = call.Name
				case *ast.SelectorExpr:

					if call.Sel.Name != "Eventually" && call.Sel.Name != "Consistently" {
						return true
					}
					callName = call.Sel.Name
				}

				anonF, err := firstArgAsFuncLit(x)
				if err != nil {
					// if the first arg is not a function then this is
					// a call with value that we can ignore here
					// https://onsi.github.io/gomega/#category-1-making-codeeventuallycode-assertions-on-values
					return true
				}

				gomegaArgName, _ := gomegaArgName(anonF, gomegaImportName)

				checkExpectCalls(pass, anonF.Body, gomegaArgName, callName)
				// need to go deeper as there can be nested Eventually blocks
				return true
			}
			return true
		})
	}
	return nil, nil
}

func firstArgAsFuncLit(f *ast.CallExpr) (*ast.FuncLit, error) {
	if len(f.Args) == 0 {
		return nil, fmt.Errorf("missing func declaration as first argument")
	}
	flit, ok := f.Args[0].(*ast.FuncLit)
	if !ok {
		return nil, fmt.Errorf("first argument is not a function but %T", f.Args[0])
	}
	return flit, nil
}

func gomegaArgName(f *ast.FuncLit, importName string) (string, error) {
	if f.Type.Params.NumFields() != 1 {
		return "", fmt.Errorf("anonymous function has no argument")
	}
	param := f.Type.Params.List[0]
	switch x := param.Type.(type) {
	case *ast.Ident:
		if importName == "" {
			if x.Name == "Gomega" {
				return param.Names[0].Name, nil
			}
		} else {
			if x.Name == importName+"."+"Gomega" {
				return param.Names[0].Name, nil
			}
		}
	case *ast.SelectorExpr:
		xId, ok := x.X.(*ast.Ident)
		if ok && xId.Name == importName && x.Sel.Name == "Gomega" {
			return param.Names[0].Name, nil
		}
	}

	return "", fmt.Errorf("anonymous function doesn't have Gomega argument")
}

func checkExpectCalls(pass *analysis.Pass, body *ast.BlockStmt, gomegaArgName string, blockName string) {
	ast.Inspect(body, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.CallExpr:
			switch call := x.Fun.(type) {
			case *ast.Ident: // We have a simple name as the call
				if call.Name == "Eventually" || call.Name == "Consistently" {
					// stop checking deeper based on the current block as there is
					// a new nested Eventually / Consistently block. This is
					// handled by the caller.
					return false
				}
				if call.Name == "Expect" {
					// We have an expect call without a package name selector
					if gomegaArgName == "" {
						tools.Report(
							pass, call.Pos(), Name,
							"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Change the "+
								"declaration of the function passed to the '%s' block to take a parameter with type "+
								"'Gomega' and use that to call 'Expect'", blockName, blockName,
						)
					} else {
						tools.Report(
							pass, call.Pos(), Name,
							"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Use '%s.Expect'",
							blockName, gomegaArgName,
						)
					}
					return false
				}
			case *ast.SelectorExpr: // we have a <name>.<funcname> expression
				if call.Sel.Name == "Eventually" || call.Sel.Name == "Consistently" {
					// stop checking deeper based on the current block as there is
					// a new nested Eventually / Consistently block. This is
					// handled by the caller.
					return false
				}
				if call.Sel.Name != "Expect" {
					return true
				}
				// we have a <name>.Expect() expression
				xId, ok := call.X.(*ast.Ident)
				if ok && xId.Name != gomegaArgName {
					if gomegaArgName == "" {
						tools.Report(
							pass, call.Pos(), Name,
							"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Change the "+
								"declaration of the function passed to the '%s' block to take a parameter with type "+
								"'Gomega' and use that to call 'Expect'", blockName, blockName,
						)
					} else {
						tools.Report(
							pass, call.Pos(), Name,
							"'Expect' in the '%s' block should be called via a local 'Gomega' parameter. Use '%s.Expect'",
							blockName, gomegaArgName,
						)
					}
					return false
				}
			}
		}
		return true
	})
}
