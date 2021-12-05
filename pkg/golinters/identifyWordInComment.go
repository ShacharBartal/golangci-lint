package golinters

import (
	identifyWordInComment "github.com/ShacharBartal/identifyWordInComment/pkg/analyzer"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func NewIdentifyWordInComment() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"identifyWordInComment",
		"Detects word(s) in comments",
		[]*analysis.Analyzer{identifyWordInComment.Analyzer},
		nil,
	)
}
