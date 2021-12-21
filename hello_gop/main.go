package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/gopbuild"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
println "Hello, World"
`

func main() {
	_, err := gossa.RunFile("main.gop", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
