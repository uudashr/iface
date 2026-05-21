package unusedmethod_test

import (
	"testing"

	"github.com/uudashr/iface/unusedmethod"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unusedmethod.Analyzer, "a")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "b")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "c")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "d")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "e")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "f")
}

func TestExclusion(t *testing.T) {
	err := unusedmethod.Analyzer.Flags.Set("exclude", "g")
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unusedmethod.Analyzer, "a")
	analysistest.Run(t, testdata, unusedmethod.Analyzer, "g")
}
