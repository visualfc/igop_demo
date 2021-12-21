package main

import (
	"go/parser"
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
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", source, parser.ParseComments)
	if err != nil {
		log.Panicln("parse", err)
	}
	ctx := gossa.NewContext(0)
	pkg, err := ctx.LoadAstFile(fset, f)
	if err != nil {
		log.Panicln("load", err)
	}
	_, err = ctx.RunPkg(pkg, "main.go", nil)
	if err != nil {
		log.Panicln("run", err)
	}
}
