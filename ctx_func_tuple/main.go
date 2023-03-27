package main

import (
	"fmt"
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package bar

func call(i, j int) (int,int) {
	return i+j, i*j
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
	v, err := interp.RunFunc("call", 100, 200)
	fmt.Println(v, err)
	if t, ok := v.(igop.Tuple); ok {
		fmt.Println(len(t), t[0], t[1])
	}
}
