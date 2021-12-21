package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package main

import "fmt"

type M struct { N int}
type T struct { N int}
func (t T) String() string {
	return "T:String" 
}
type U struct {
	T
}
func (u U) GoString() string {
	return "U:GoString"
}

func main() {
	m := &M{100}
	t := &T{100}
	u := &U{T{100}}
	fmt.Printf("%v %#v\n",m,m)
	fmt.Printf("%v %#v\n",t,t)
	fmt.Printf("%v %#v\n",u,u)
}
`

func main() {
	_, err := gossa.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
