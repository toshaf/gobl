package install

import (
	"archive/tar"
	"compress/gzip"
	"github.com/toshaf/gobl/cmd"
	"github.com/toshaf/gobl/log"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func NewInstallCmd(logger log.Logger, fname string) cmd.Cmd {
	return &installT{
		Logger: logger,
		fname:  fname,
	}
}

type installT struct {
	log.Logger
	fname string
}

func (ins *installT) Run() error {
	in, err := os.Open(ins.fname)
	if err != nil {
		return err
	}
	defer in.Close()

	gr, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	rdr := tar.NewReader(gr)

	atime := time.Now()
	// time since epoch stored in signed 32 bit integers wrap around
	// on January 19 2038 at 03:14:07 -- stopping just short of that
	mtime := time.Date(2038, time.January, 19, 0, 0, 0, 0, time.UTC)

	for {
		header, err := rdr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		ins.Logv("Extracting %s", header.Name)

		fname := path.Join(os.ExpandEnv("$GOPATH"), header.Name)

		ins.Logv("Ensuring %s exists", filepath.Dir(fname))
		os.MkdirAll(filepath.Dir(fname), os.ModeDir|os.ModePerm)

		out, err := os.Create(fname)
		if err != nil {
			ins.Logv("ERROR creating %s: %s", fname, err)
			return err
		}

		if _, err = io.Copy(out, rdr); err != nil {
			ins.Logv("ERROR: %s", err)
			return err
		}

		out.Close()
		if strings.HasSuffix(fname, ".a") {
			// make sure .a files look newer than
			// the source files so Go doesn't try
			// to rebuild them
			os.Chtimes(fname, atime, mtime)
		}

		ins.Logv("Done with %s", header.Name)
	}

	return nil
}
