package main

import (
	"fmt"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	panic("error")
}
`

func main() {
	_, err := igop.RunFile("main.go", source, nil, 0)
	if err != nil {
		fmt.Println("PANIC:", err)
	}
}
