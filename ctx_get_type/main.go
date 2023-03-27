package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
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
	ctx := igop.NewContext(0)
	pkg, err := ctx.LoadFile("main.go", source)
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
