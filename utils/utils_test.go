package utils_test

import (
	"github.com/toshaf/gobl/utils"
	"strings"
	"testing"
)

func FakeExpand(v string) string {
	return strings.Replace(v, "$GOPATH", "/usr/tosh/code/go", -1)
}

func Test_Paths(t *testing.T) {
	paths := utils.NewPaths("github.com/toshaf/exhibit", FakeExpand)

	if paths.PkgName != "github.com/toshaf/exhibit" {
		t.Errorf("Wrong PkgName: %s", paths.PkgName)
	}

	if paths.GoblPkgDir != "/usr/tosh/code/go/gobl-pkg/" {
		t.Errorf("Wrong GoblPkgDir: %s", paths.GoblPkgDir)
	}

	if paths.GoblFilename != "/usr/tosh/code/go/gobl-pkg/github.com.toshaf.exhibit.gobl" {
		t.Errorf("Wrong GoblFilename: %s", paths.GoblFilename)
	}
}
