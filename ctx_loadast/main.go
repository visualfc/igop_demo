package main

import (
	"go/parser"
	"go/token"
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
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", source, parser.ParseComments)
	if err != nil {
		log.Panicln("parse", err)
	}
	ctx := igop.NewContext(0)
	ctx.FileSet = fset
	pkg, err := ctx.LoadAstFile("main.go", f)
	if err != nil {
		log.Panicln("load", err)
	}
	_, err = ctx.RunPkg(pkg, "main.go", nil)
	if err != nil {
		log.Panicln("run", err)
	}
}
