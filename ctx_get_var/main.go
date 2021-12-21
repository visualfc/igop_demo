package main

import (
	"fmt"
	"go/token"
	"log"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package bar

var index int = 100
var pindex *int = &index

func show() {
	println(index)
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
			fmt.Println(*p)
			*p = 200
		}
	}
	if v, ok := interp.GetVarAddr("pindex"); ok {
		if p, ok := v.(**int); ok {
			fmt.Println(**p)
			**p = 300
		}
	}
	if fn, ok := interp.GetFunc("show"); ok {
		fn.(func())()
	}
}
