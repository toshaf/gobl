package pack

import (
	"archive/tar"
	"fmt"
	"github.com/toshaf/gobl/source"
	"github.com/toshaf/gobl/utils"
	"io/ioutil"
	"path"
	"strings"
)

func (c *PackCmd) packSource(writer *tar.Writer, p *utils.Paths) error {
	base := path.Join(p.SrcDir, p.PkgName)
	files, err := ioutil.ReadDir(base)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			fpath := path.Join(base, file.Name())
			src, err := ioutil.ReadFile(fpath)
			if err != nil {
				return err
			}

			src, err = source.Prune(src)

			header := tar.Header{
				Name: fmt.Sprintf("src/%s/%s", p.PkgName, file.Name()),
				Mode: 0444,
				Size: int64(len(src)),
			}

			writer.WriteHeader(&header)
			writer.Write(src)
		}
	}

	return nil
}
