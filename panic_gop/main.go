package main

import (
	"fmt"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
var i int
println 100/i
`

func main() {
	_, err := igop.RunFile("main.gop", source, nil, 0)
	if err != nil {
		fmt.Println("PANIC:", err)
	}
}
