package source_test

import (
	. "github.com/toshaf/exhibit"
	"github.com/toshaf/gobl/source"
	"io/ioutil"
	"testing"
)

func Test_Prune(t *testing.T) {
	src, err := ioutil.ReadFile("../test-files/a/a.go")
	check(err, t)

	pruned, err := source.Prune(src)
	check(err, t)

	Exhibit.A(Text(pruned), t)
}

func check(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
