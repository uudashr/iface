package opaque_test

import (
	"testing"

	"github.com/uudashr/iface/opaque"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "a")
	analysistest.Run(t, testdata, opaque.Analyzer, "b")
	analysistest.Run(t, testdata, opaque.Analyzer, "c")
	analysistest.Run(t, testdata, opaque.Analyzer, "d")
	analysistest.Run(t, testdata, opaque.Analyzer, "e")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "f")
	analysistest.Run(t, testdata, opaque.Analyzer, "g")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "h")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "i")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "j")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "k")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "l")
	analysistest.Run(t, testdata, opaque.Analyzer, "m")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "n")
	analysistest.Run(t, testdata, opaque.Analyzer, "o")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "p")
	analysistest.RunWithSuggestedFixes(t, testdata, opaque.Analyzer, "z")
}
