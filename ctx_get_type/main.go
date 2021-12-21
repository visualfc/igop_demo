package main

import (
	"fmt"
	"go/token"
	"log"
	"reflect"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package bar

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) Set(x, y int) {
	p.x,p.y = x,y
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v,%v)",p.x,p.y)
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
	if typ, ok := interp.GetType("Point"); ok {
		v := reflect.New(typ)
		fn := v.MethodByName("Set")
		fn.Interface().(func(int, int))(100, 200)
		fmt.Println(v)
	}
}
