package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
)

var source = `
package main

import "fmt"

type T struct {}

func (t T) String() string {
	return "Hello, World" 
}

func main() {
	fmt.Println(&T{})
}
`

func main() {
	_, err := gossa.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
