package main

import (
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
	_, err := igop.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
