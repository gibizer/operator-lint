package main

import (
	"github.com/gibizer/envtest-linter/linters/L001"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		L001.LocalGomegaEventuallyAnalyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
