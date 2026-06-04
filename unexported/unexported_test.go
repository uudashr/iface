package unexported_test

import (
	"testing"

	"github.com/uudashr/iface/unexported"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unexported.Analyzer,
		"fun",
		"method",
		"ignoredirective",
		"typeparams")
}
