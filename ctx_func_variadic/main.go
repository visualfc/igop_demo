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

func sum(n ...int) (r int) {
	for _, v := range n {
		r += v
	}
	return
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
	v1, err := interp.RunFunc("sum", []int{100, 200})
	fmt.Println(v1, err)
	v2, err := interp.RunFunc("sum", []int{})
	fmt.Println(v2, err)
	// error call
	v3, err := interp.RunFunc("sum")
	fmt.Println(v3, err)
}
