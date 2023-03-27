package main

import (
	"fmt"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/os"
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
	code, err := igop.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("exit", code)
}
