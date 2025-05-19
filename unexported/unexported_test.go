package unexported_test

import (
	"testing"

	"github.com/uudashr/iface/unexported"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unexported.Analyzer, "a")
	analysistest.Run(t, testdata, unexported.Analyzer, "b")
}
