package gocoder

import (
	"bufio"
	"fmt"
	"io"
)

type SourceCode struct {
	Package Package
	Imports DefineGroup
	Items   []fmt.Stringer
}

func NewSourceFile() *SourceCode {
	sf := new(SourceCode)
	sf.Imports.Type = "import"
	return sf
}

func (sf *SourceCode) AddImport(i *Import) {
	sf.Imports.Items = append(sf.Imports.Items, i)
}

func (sf *SourceCode) ImportAs(path, name string) {
	if !sf.isImported(path, name) {
		sf.AddImport(&Import{
			Path: path,
			Name: name,
		})
	}
}

func (sf *SourceCode) isImported(path, name string) bool {
	for _, i := range sf.Imports.Items {
		imp := i.(*Import)
		if imp.Path == path && imp.Name == name {
			return true
		}
	}
	return false
}

func (sf *SourceCode) Import(path string) {
	sf.ImportAs(path, "")
}

// add item
func (sf *SourceCode) Add(item fmt.Stringer) {
	sf.Items = append(sf.Items, item)
}

// add const group
func (sf *SourceCode) AddConstGroup(args ...string) *DefineGroup {
	nargs := len(args)
	if !(nargs == 0 || nargs == 1) {
		panic("length of args must be 0 or 1")
	}
	var desc string
	if nargs == 1 {
		desc = args[0]
	}
	group := &DefineGroup{
		Desc: desc,
		Type: "const",
	}

	sf.Add(group)
	return group
}

// TODO add var group

func (sf *SourceCode) Line(l string) {
	sf.Add(String(l + "\n"))
}

func (sf *SourceCode) Linef(format string, a ...interface{}) {
	sf.Add(String(fmt.Sprintf(format+"\n", a...)))
}

func (sf *SourceCode) WriteTo(w io.Writer) {
	writer := bufio.NewWriter(w)
	// package header
	writer.WriteString(sf.Package.String())
	// imports
	writer.WriteString(sf.Imports.String())

	for _, item := range sf.Items {
		writer.WriteString(item.String())
	}

	writer.Flush()
}

// TODO sf.SaveToFile
