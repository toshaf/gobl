# GOBL - The GO Binary Library
The aim of GOBL is to provide a means for the distribution of GO packages without necessarily sharing source code. These packages should still support normal GO tooling for build and documentation.

## GOBL package files
GOBL files are intended to be a flexible for container for binary + partial source or source only distribution of GO code.

GOBL files are simply Gzip compressed tarballs with a .gobl extension containing a subset of a GO work area; as such you can extract the contents of one by issuing

    $ tar -xzf some.package.gobl

An example GOBL file would have the following structure

    src/
        github.com/
            banana/
                split.go
    pkg/
        linux_amd64/
            banana/
                split.a

The immediate child dirs of pkg/ have the form `$GOOS_$GOARCH`

## The `gobl` command line tool
The `gobl` command line tool can be used to pack, publish and consume GOBL packages.

## Installation
The `gobl` tool is installed from source using

    $ go get github.com/toshaf/gobl

Just make sure it's on your $PATH` and you're ready to go.

### General usage
The general usage pattern for the `gobl` tool is

    $ gobl <command> <options> <inputs>

Where `command` specifies the action you wish to take, `inputs` are the inputs specific to the command (if any) and `options` are the set of adjustments/configurations you wish to specify, again specific to the chosen command.

#### Environment variables
The `gobl` tool requires the `$GOPATH` environment variable to be set - all package paths are based on this.

### Commands
#### `pack`
The `pack` command will create a GOBL package from the GO package(s) specified.

The name of the file generated will be the longest common prefix of all the packages packed with forward slashes converted to dots and with the .gobl extension; packing `github.com/toshaf/exhibit` would result in a GOBL package called `github.com.toshaf.exhibit.gobl`.

A single input specifying name of the GO package(s) to be `pack`ed is required; this takes the form of either a single package name or the name of a package to recursively pack.

To pack the package `github.com/banana/split` you'd issue

    $ gobl pack github.com/banana/split

To pack all packages under `github.com/banana/split/` you'd issue

    $ gobl pack github.com/banana/split/...

Either way you'd end up with a file called `$GOPATH/gobl-pkg/github.com.banana.split.gobl`.

#### `install`
The `install` command unpacks a GOBL file for consumption by GO programs placing the contents under `$GOPATH/src` and `$GOPATH/pkg` as appropriate.

The `gobl` tool takes a single input specifying the full path to the GOBL file to be installed.

Note: to enable consumers to build against binary distributions GOBL sets the mtimes of the .a files such that they look newer to GO than their counterpart .go files - you are discouraged from altering the mtimes of these files. Build errors resulting from this change can be fixed by simply re-installing a GOBL file.
