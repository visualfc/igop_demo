package main

import (
	"fmt"
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/math"
)

var source = `
package bar

import "math"

const Pi = math.Pi

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
	if v, ok := interp.GetConst("Pi"); ok {
		fmt.Println(v)
		fmt.Println(v.ExactString())
	}
}
