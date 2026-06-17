package unusedmethod_test

import (
	"testing"

	"github.com/uudashr/iface/unusedmethod"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	err := unusedmethod.Analyzer.Flags.Set("exclude", "excludepkg")
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unusedmethod.Analyzer,
		"basic")

	analysistest.Run(t, testdata, unusedmethod.Analyzer,
		"excludepkg",
		"methodref",
		"methodexpr")
}
