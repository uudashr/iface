package main

import (
	"github.com/uudashr/iface/duplicate"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(duplicate.Analyzer)
}
