package unused_test

import (
	"testing"

	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "a")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "agroup")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "agroupdoc")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "agroupsingle")
	analysistest.Run(t, testdata, unused.Analyzer, "b")
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "d")
	analysistest.Run(t, testdata, unused.Analyzer, "e")
}

func TestExclusion(t *testing.T) {
	err := unused.Analyzer.Flags.Set("exclude", "c")
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer, "a")
	analysistest.Run(t, testdata, unused.Analyzer, "c")
}
