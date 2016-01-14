package main

import (
	"fmt"
	"github.com/toshaf/gobl/cli"
	"github.com/toshaf/gobl/cmd/install"
	"github.com/toshaf/gobl/cmd/pack"
	"github.com/toshaf/gobl/log"
	"github.com/toshaf/gobl/utils"
	"os"
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

	args := cli.Parse()
	if len(args.Inputs) == 0 {
		Panic("Usage: %s <cmd> <args>\n", args.Cmd)
	}

	logger := log.NewLoggerFromFlags()

	switch args.Cmd {
	case "pack":
		if len(args.Inputs) < 1 {
			Panic("Usage: %s pack <pkg-path>\n", args.Cmd)
		}

		pkgname := args.Inputs[0]
		paths := utils.NewPaths(pkgname)

		outfile, err := paths.CreateOutputFile()
		if err != nil {
			panic(err)
		}
		defer outfile.Close()

		logger.Logv("Packing %s", pkgname)

		packer := pack.NewPackCmd(logger, paths, outfile)
		err = packer.Run()
		if err != nil {
			panic(err)
		}

		logger.Logv("Written %s", paths.GoblFilename)
	case "install":
		if len(args.Inputs) < 1 {
			Panic("Usage: %s install <gobl-file>\n", args.Cmd)
		}

		fname := args.Inputs[0]
		logger.Logv("Installing %s", fname)

		installer := install.NewInstallCmd(logger, fname)
		err := installer.Run()
		if err != nil {
			panic(err)
		}

		logger.Logv("Installed %s", fname)
	default:
		panic("Unknown cmd: " + args.Cmd)
	}
}

func Panic(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}
