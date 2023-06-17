package base

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
)

type FileLinter interface {
	LintFile(*ast.File) error
}

type BaseLinter struct {
	FileLinter
	*analysis.Analyzer
	Pass *analysis.Pass
	log  *log.Logger
	skip *bool
}

func NewBaseLinter(Name string, Doc string, l FileLinter) *BaseLinter {
	base := &BaseLinter{
		Analyzer: &analysis.Analyzer{
			Name: Name,
			Doc:  Doc,
		},
		log: log.New(os.Stderr, Name+": ", log.Lshortfile|log.Lmsgprefix),
	}
	base.FileLinter = l
	base.Analyzer.Run = base.baseRun

	base.skip = base.Analyzer.Flags.Bool(
		"skip", false, "Use to skip the given check")

	return base
}

func (l *BaseLinter) baseRun(pass *analysis.Pass) (interface{}, error) {
	if *l.skip {
		return nil, nil
	}

	l.Pass = pass
	for _, file := range l.Pass.Files {
		err := l.LintFile(file)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// Report a lint rule violation
func Report(pass *analysis.Pass, pos token.Pos, analyzerName string, format string, a ...any) {
	pass.Reportf(pos, analyzerName+": "+format, a...)
}

// Report a lint rule violation
func (l *BaseLinter) Report(pos token.Pos, format string, a ...any) {
	l.Pass.Reportf(pos, l.Name+": "+format, a...)
}

func (l *BaseLinter) Debug(pos token.Pos, format string, a ...any) {
	tokenPos := l.Pass.Fset.Position(pos)
	srcAbsPath := tokenPos.Filename
	srcPath, err := filepath.Rel(os.Getenv("PWD"), srcAbsPath)
	if err != nil {
		srcPath = srcAbsPath
	}
	path := fmt.Sprintf("on %s:%d:%d: ", srcPath, tokenPos.Line, tokenPos.Column)
	l.log.Printf(path+format, a...)
}
