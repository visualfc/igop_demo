package main

import (
	"log"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
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
	_, err = interp.RunMain()
	if err != nil {
		log.Panicln("run", err)
	}
}
