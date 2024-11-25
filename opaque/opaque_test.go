package opaque_test

import (
	"testing"

	"github.com/uudashr/iface/opaque"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "a")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "b")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "c")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "d")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "e")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "f")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "g")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "h")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "x")
}
