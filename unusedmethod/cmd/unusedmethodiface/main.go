package main

import (
	"github.com/uudashr/iface/unusedmethod"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unusedmethod.Analyzer)
}
