package main

import (
	"fmt"
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package bar

func add(i, j int) int {
	return i+j
}
`

func main() {
	ctx := igop.NewContext(0)
	pkg, err := ctx.LoadFile("main.go", source)
	if err != nil {
		log.Panicln("load", err)
	}
	interp, err := ctx.NewInterp(pkg)
	if err != nil {
		log.Panicln("interp", err)
	}
	if v, ok := interp.GetFunc("add"); ok {
		if fn, ok := v.(func(int, int) int); ok {
			r := fn(100, 200)
			fmt.Println(r)
		}
	}
}
