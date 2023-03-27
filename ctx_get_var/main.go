package main

import (
	"fmt"
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
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
	ctx := igop.NewContext(0)
	pkg, err := ctx.LoadFile("main.go", source)
	if err != nil {
		log.Panicln("load", err)
	}
	interp, err := ctx.NewInterp(pkg)
	if err != nil {
		log.Panicln("interp", err)
	}
	if err := interp.RunInit(); err != nil {
		log.Panicln("init", err)
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
