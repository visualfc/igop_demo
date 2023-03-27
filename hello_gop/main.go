package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
println "Hello, World"
`

func main() {
	_, err := igop.RunFile("main.gop", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
