package empty_test

import (
	"testing"

	"github.com/uudashr/iface/empty"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, empty.Analyzer, "a")
}
