package main

import (
	"github.com/uudashr/iface/duplicate"
	"github.com/uudashr/iface/opaque"
	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		unused.Analyzer,
		duplicate.Analyzer,
		opaque.Analyzer,
	)
}
