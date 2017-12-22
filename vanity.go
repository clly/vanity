package main

// +build linux

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/clly/vanity/static"
	"github.com/pkg/errors"
)

type vanity struct {
	Package      string
	ShortPackage string
	Name         string
}

//go:generate petrify -pkg static -o static/bindata.go -ignore=bindata.go static/...
func main() {
	var v vanity

	dirs, err := parse(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	v.Package = filepath.Join(dirs...)
	v.ShortPackage = filepath.Join(dirs[1:]...)
	v.Name = dirs[len(dirs)-1]

	err = vanityLabels(os.Stdout, v)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	writer, err := getWriter(".travis.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer writer.Close()
	err = writeTravis(writer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func parse(args []string) ([]string, error) {
	var dirs []string
	if len(args) != 2 {
		return dirs, fmt.Errorf("Too many or not enough args %s: %d args", os.Args[1:], len(os.Args))
	}

	dirs = strings.Split(args[1], "/")

	if len(dirs) <= 1 {
		return dirs, errors.New(fmt.Sprint("Can't have a package with only 1 separator ", dirs))
	}

	return dirs, nil
}

func vanityLabels(writer io.Writer, v vanity) error {
	t := static.MustParseTemplates("static/vanity.tmpl")
	return t.Execute(writer, v)
}

func writeTravis(writer io.Writer) error {
	t := static.MustParseTemplates("static/travis.tmpl")
	return t.Execute(writer, nil)
}

func getWriter(location string) (io.WriteCloser, error) {
	f, err := os.OpenFile(location, os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}
