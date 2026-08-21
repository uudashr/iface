package main

func main() {
	var g G
	var _ I = g

	g.m("foo")

	var h H
	h.m("foo")
}

type G struct{}

func (G) m(string) {}

type I interface {
	m(string) // want "^method 'm\\(\\)' is declared on interface 'I' but not used within the package$"
}

type H struct{}

func (H) m(string) {}
