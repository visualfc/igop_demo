package main

import (
	"go/token"
	"log"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package bar

func call(i, j int) (int,int) {
	return i+j, j*j
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
	v, err := interp.RunFunc("call", 100, 200)
	log.Println(v, err)
	t := v.(gossa.Tuple)
	log.Println(len(t), t[0], t[1])
}
