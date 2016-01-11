package source

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
)

func Prune(src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	ast.FileExports(file)

	pruner := NewPruner()
	ast.Walk(pruner, file)

	file.Comments = nil

	buf := bytes.Buffer{}
	err = printer.Fprint(&buf, fset, file)
	if err != nil {
		return nil, err
	}

	pruned, err := ioutil.ReadAll(&buf)
	return pruned, err
}

type Pruner struct {
}

func NewPruner() ast.Visitor {
	return &Pruner{}
}

func (p *Pruner) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	if fn, ok := node.(*ast.FuncDecl); ok {
		fn.Body = nil
		return nil
	}

	return p
}
