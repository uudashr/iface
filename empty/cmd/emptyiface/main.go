package main

import (
	"github.com/uudashr/iface/empty"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(empty.Analyzer)
}
