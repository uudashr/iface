package opaque_test

import (
	"testing"

	"github.com/uudashr/iface/opaque"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, opaque.Analyzer, "a")
	analysistest.Run(t, testdata, opaque.Analyzer, "b")
	analysistest.Run(t, testdata, opaque.Analyzer, "c")
	analysistest.Run(t, testdata, opaque.Analyzer, "d")
	analysistest.Run(t, testdata, opaque.Analyzer, "e")
	analysistest.Run(t, testdata, opaque.Analyzer, "x")
}
