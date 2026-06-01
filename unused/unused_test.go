package unused_test

import (
	"testing"

	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	err := unused.Analyzer.Flags.Set("exclude", "excludepkg")
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, unused.Analyzer,
		"basic",
		"ignoredirective",
		"excludepkg")
}
