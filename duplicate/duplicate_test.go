package duplicate_test

import (
	"testing"

	"github.com/uudashr/iface/duplicate"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, duplicate.Analyzer, "a")
	analysistest.Run(t, testdata, duplicate.Analyzer, "b")
	analysistest.Run(t, testdata, duplicate.Analyzer, "c")
	analysistest.Run(t, testdata, duplicate.Analyzer, "d")
}
