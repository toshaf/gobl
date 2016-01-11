package main

import (
	"fmt"
	"github.com/toshaf/gobl/cmd/pack"
	"os"
	"path"
	"strings"
)

func main() {
	defer func() {
		p := recover()
		if p != nil {
			if err, ok := p.(error); ok {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Println(p)
			os.Exit(1)
		}
	}()

	prog := os.Args[0]

	if len(os.Args) < 2 {
		Panic("Usage: %s <cmd> <args>\n", prog)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "pack":
		if len(args) < 1 {
			Panic("Usage: %s pack <pkg-path>\n", prog)
		}

		goblpkg := os.ExpandEnv("$GOPATH/gobl-pkg/")
		err := os.MkdirAll(goblpkg, os.ModeDir|os.ModePerm)
		if err != nil {
			panic(err)
		}

		basename := args[0]

		outfilename := path.Join(goblpkg, strings.Replace(basename, "/", ".", -1)) + ".gobl"

		outfile, err := os.Create(outfilename)
		if err != nil {
			panic(err)
		}
		defer outfile.Close()

		err = pack.Pack(basename, outfile)
		if err != nil {
			panic(err)
		}
	}
}

func Panic(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}
