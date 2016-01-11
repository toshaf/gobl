package cli

import (
	"os"
	"strings"
)

type Args struct {
	Exec    string
	Cmd     string
	Options []string
	Inputs  []string
}

func Parse() Args {
	return ParseArgs(os.Args)
}

func ParseArgs(args []string) Args {
	exec := ""
	if len(args) > 0 {
		exec = args[0]
		args = args[1:]
	}

	options := []string{}
	inputs := []string{}
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			options = append(options, a)
		} else {
			inputs = append(inputs, a)
		}
	}

	cmd := ""
	if len(inputs) > 0 {
		cmd = inputs[0]
		inputs = inputs[1:]
	}

	return Args{
		Exec:    exec,
		Cmd:     cmd,
		Options: options,
		Inputs:  inputs,
	}
}
