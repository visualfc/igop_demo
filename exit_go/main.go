package main

import (
	"fmt"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/os"
)

var source = `
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World")
	os.Exit(2)
}
`

func main() {
	code, err := gossa.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("exit", code)
}
