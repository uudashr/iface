package main

import (
	"github.com/uudashr/iface/opaque"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(opaque.Analyzer)
}
