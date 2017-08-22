package main

import (
	"github.com/AstromechZA/go-multi-project-demo/aaaa"
	"github.com/AstromechZA/go-multi-project-demo/bbbb"

	"github.com/AstromechZA/go-multi-project-demo/dep/sub"
)

func main() {
	aaaa.A()
	bbbb.B()
	sub.Sub()
}
