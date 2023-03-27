package main

import (
	"os"
	"path/filepath"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg"
	_ "github.com/goplus/ipkg/golang.org/x/image/vector"
	_ "github.com/goplus/ipkg/golang.org/x/sys/unix"
	_ "github.com/goplus/reflectx/icall/icall65536"
)

func main() {
	root, args := os.Args[1], os.Args[2:]
	if !filepath.IsAbs(root) {
		wd, _ := os.Getwd()
		root = filepath.Join(wd, root)
	}
	code, err := igop.Run(root, args, igop.EnableDumpImports)
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}
