package main

import (
	"github.com/uudashr/iface/identical"
	"github.com/uudashr/iface/opaque"
	"github.com/uudashr/iface/unexported"
	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		unused.Analyzer,
		identical.Analyzer,
		opaque.Analyzer,
		unexported.Analyzer,
	)
}
