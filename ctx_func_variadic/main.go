package main

import (
	"fmt"
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
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
	ctx := igop.NewContext(0)
	pkg, err := ctx.LoadFile("main.go", source)
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
}
