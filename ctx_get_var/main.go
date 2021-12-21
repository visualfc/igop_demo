package main

import (
	"go/token"
	"log"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package bar

var index int = 100

func add(i, j int) int {
	return i+j
}
`

func main() {
	fset := token.NewFileSet()
	ctx := gossa.NewContext(0)
	pkg, err := ctx.LoadFile(fset, "main.go", source)
	if err != nil {
		log.Panicln("load", err)
	}
	interp, err := ctx.NewInterp(pkg)
	if err != nil {
		log.Panicln("interp", err)
	}
	if v, ok := interp.GetVarAddr("index"); ok {
		if p, ok := v.(*int); ok {
			log.Println(*p)
		}
	}
}
