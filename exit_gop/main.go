package main

import (
	"fmt"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/pkg/os"
)

var source = `
import "os"
println "Hello, World"
os.exit(2)
`

func main() {
	code, err := igop.RunFile("main.gop", source, nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("exit code", code)
}
