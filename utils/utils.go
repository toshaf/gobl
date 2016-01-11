package utils

import (
	"compress/gzip"
	"io"
	"os"
	"path"
	"strings"
)

type Paths struct {
	PkgName      string
	GoblPkgDir   string
	GoblFilename string
	SrcDir       string
}

type EnvExpander func(string) string

type P []EnvExpander

func (params P) Expand(v string) string {
	if len(params) > 0 {
		return params[0](v)
	}

	return os.ExpandEnv(v)
}

func NewPaths(pkgname string, params ...EnvExpander) *Paths {
	ps := P(params)
	goblpkg := ps.Expand("$GOPATH/gobl-pkg/")
	fname := path.Join(goblpkg, strings.Replace(pkgname, "/", ".", -1)) + ".gobl"
	src := ps.Expand("$GOPATH/src")
	return &Paths{
		PkgName:      pkgname,
		GoblPkgDir:   goblpkg,
		GoblFilename: fname,
		SrcDir:       src,
	}
}

func (p *Paths) CreateOutputFile() (io.WriteCloser, error) {
	err := os.MkdirAll(p.GoblPkgDir, os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(p.GoblFilename)
	if err != nil {
		return nil, err
	}

	zip := gzip.NewWriter(file)

	return zip, err
}
