package main

import (
	"github.com/uudashr/iface/unexported"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unexported.Analyzer)
}
