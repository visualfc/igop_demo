package main

import (
	"fmt"
	"go/token"
	"log"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}

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
	v, err := interp.RunFunc("add", 100, 200)
	fmt.Println(v, err)
}
