package main

import (
	linters "github.com/gibizer/envtest-linter"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.TodoAnalyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
