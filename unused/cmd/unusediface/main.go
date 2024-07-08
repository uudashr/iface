package main

import (
	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unused.Analyzer)
}
