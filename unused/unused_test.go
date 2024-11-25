package unused_test

import (
	"testing"

	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	err := unused.Analyzer.Flags.Set("exclude", "c")
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "a")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "agroup")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "b")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "c")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "d")
}
