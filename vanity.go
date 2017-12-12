package main

// +build linux

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/clly/vanity/static"
	"strings"
)

type vanity struct {
	Package string
	ShortPackage string
}

//go:generate petrify -debug -pkg static -o static/bindata.go -ignore=bindata.go static/...
func main() {
	var v vanity
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Too many args")
	}

	dirs := strings.Split(os.Args[1], "/")

	if len(dirs) < 1 {
		fmt.Fprintln(os.Stderr,  "Can't have a package with only 1 separator ", dirs)
	}
	v.Package = filepath.Join(dirs...)
	v.ShortPackage = filepath.Join(dirs[1:]...)

	t := static.MustParseTemplates("static/vanity.tmpl")
	t.Execute(os.Stdout, v)
}