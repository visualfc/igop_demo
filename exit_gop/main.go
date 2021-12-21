package main

import (
	"fmt"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/gopbuild"
	_ "github.com/goplus/gossa/pkg/os"
)

var source = `
import "os"
println "Hello, World"
os.exit(2)
`

func main() {
	code, err := gossa.RunFile("main.gop", source, nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("exit code", code)
}
