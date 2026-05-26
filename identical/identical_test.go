package identical_test

import (
	"testing"

	"github.com/uudashr/iface/identical"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, identical.Analyzer,
		"basic",
		"embedded",
		"typeconstraints",
		"typedefalias",
		"ignoredirective")
}
