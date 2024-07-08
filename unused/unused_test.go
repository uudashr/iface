package unused_test

import (
	"testing"

	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unused.Analyzer, "a")
	analysistest.Run(t, testdata, unused.Analyzer, "b")
}
