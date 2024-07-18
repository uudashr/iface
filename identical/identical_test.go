package identical_test

import (
	"testing"

	"github.com/uudashr/iface/identical"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, identical.Analyzer, "a")
	analysistest.Run(t, testdata, identical.Analyzer, "b")
	analysistest.Run(t, testdata, identical.Analyzer, "c")
	analysistest.Run(t, testdata, identical.Analyzer, "d")
	analysistest.Run(t, testdata, identical.Analyzer, "e")
}
