package pack

import (
	"archive/tar"
	"fmt"
	"github.com/toshaf/gobl/cmd"
	"github.com/toshaf/gobl/log"
	"github.com/toshaf/gobl/utils"
	"io"
	"io/ioutil"
	"os"
)

const (
	GO_LIB_EXT = ".a"
)

func NewPackCmd(logger log.Logger, paths *utils.Paths, writer io.Writer) *PackCmd {
	return &PackCmd{
		log:    logger,
		writer: writer,
		Paths:  paths,
	}
}

type PackCmd struct {
	log    log.Logger
	writer io.Writer
	Paths  *utils.Paths
}

var _ cmd.Cmd = &PackCmd{}

func (c *PackCmd) Run() error {
	pkgdir := os.ExpandEnv("$GOPATH/pkg")
	c.log.Logv("Package dir: %s", pkgdir)
	files, err := ioutil.ReadDir(pkgdir)
	if err != nil {
		return err
	}

	writer := tar.NewWriter(c.writer)
	defer func() {
		if err := writer.Close(); err != nil {
			panic(err)
		}
	}()

	for _, f := range files {
		if f.IsDir() {
			c.log.Logv("OS/Arch: %s", f.Name())
			if err := c.packLib(writer, f.Name()); err == nil {
				if err = c.packSource(writer, c.Paths); err != nil {
					panic(err)
				}
			} else if os.IsNotExist(err) {
				c.log.Logv("not built")
				continue // it might not be built for this combo
			} else {
				panic(err)
			}
		}
	}

	return nil
}

func (c *PackCmd) packLib(writer *tar.Writer, dir string) error {
	// these should be of the form $GOOS_$GOARCH
	name := fmt.Sprintf("pkg/%s/%s%s", dir, c.Paths.PkgName, GO_LIB_EXT)
	fname := os.ExpandEnv("$GOPATH/") + name
	c.log.Logv("Looking for %s", fname)
	file, err := os.Open(fname)
	if err != nil {
		return err
	}

	defer file.Close()

	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	header := tar.Header{
		Name: name,
		Mode: 0444,
		Size: int64(len(body)),
	}

	writer.WriteHeader(&header)
	writer.Write(body)

	return nil
}
