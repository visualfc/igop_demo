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
	_, err = ctx.RunPkg(pkg, "main.go", nil)
	if err != nil {
		log.Panicln("run", err)
	}
}
