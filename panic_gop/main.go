package main

import (
	"fmt"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/gopbuild"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
var i int
println 100/i
`

func main() {
	_, err := gossa.RunFile("main.gop", source, nil, 0)
	if err != nil {
		fmt.Println("PANIC:", err)
	}
}
