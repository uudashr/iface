package main

import (
	"github.com/uudashr/iface/identical"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(identical.Analyzer)
}
