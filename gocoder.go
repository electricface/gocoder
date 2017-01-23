package gocoder

import (
	"fmt"
	"regexp"
	"strings"
)

func Comment(in string) string {
	if in == "" {
		return ""
	}
	return "// " + in + "\n"
}

type Package struct {
	Desc string
	Name string
}

func (p *Package) String() string {
	if p.Name == "" {
		panic("Package.Name is empty")
	}
	return Comment(p.Desc) + "package " + p.Name + "\n"
}

type DefineGroup struct {
	Desc  string
	Items []DefineItem
	Type  string // var const or import
}

func (dg *DefineGroup) Add(item DefineItem) {
	dg.Items = append(dg.Items, item)
}

func (dg *DefineGroup) String() string {
	var parts []string
	if len(dg.Items) == 0 {
		return ""
	}
	for _, item := range dg.Items {
		if dg.Type != item.Type() {
			panic("define group element type")
		}
		parts = append(parts, item.ItemString())
	}
	return Comment(dg.Desc) + dg.Type + "(\n" + strings.Join(parts, "\n") + "\n)\n"
}

type DefineItem interface {
	String() string
	ItemString() string
	Type() string
}

type Import struct {
	Desc string
	Path string
	Name string
}

func (i *Import) String() string {
	return Comment(i.Desc) + "import " + i.Name + " \"" + i.Path + "\"\n"
}

func (i *Import) ItemString() string {
	return Comment(i.Desc) + i.Name + " \"" + i.Path + "\""
}

func (_ *Import) Type() string {
	return "import"
}

type TypeAlias struct {
	Desc string
	New  string
	Old  string
}

func (ta *TypeAlias) String() string {
	return Comment(ta.Desc) + "type " + ta.New + " " + ta.Old + "\n"
}

type Const struct {
	Desc  string
	Name  string
	Typ   string
	Value interface{}
}

func (_ *Const) Type() string {
	return "const"
}

func (c *Const) String() string {
	str := Comment(c.Desc) + "const " + c.Name + " " + c.Typ + " "
	if c.Value == nil {
		return str + "\n"
	}
	return str + " = " + fmt.Sprintf("%#v", c.Value) + "\n"
}

func (c *Const) ItemString() string {
	str := Comment(c.Desc) + c.Name + " " + c.Typ + " "
	if c.Value == nil {
		return str
	}
	return str + " = " + fmt.Sprintf("%#v", c.Value)
}

type Var struct {
	Desc  string
	Name  string
	Typ   string
	Value string
}

func (_ *Var) Type() string {
	return "var"
}

func (v *Var) String() string {
	str := Comment(v.Desc) + "var " + v.Name + " " + v.Typ + " "
	if v.Value == "" {
		return str + "\n"
	}
	return str + " = " + v.Value + "\n"
}

func (v *Var) ItemString() string {
	str := Comment(v.Desc) + v.Name + " " + v.Typ + " "
	if v.Value == "" {
		return str
	}
	return str + " = " + v.Value
}

type String string

func (s String) String() string {
	return string(s)
}

// AllCaps is a regex to test if a string identifier is made of
// all upper case letters.
var allCaps = regexp.MustCompile("^[A-Z0-9]+$")

func PublicName(s string) string {
	// If the string is all caps, lower it and capitalize first letter.
	if allCaps.MatchString(s) {
		return strings.Title(strings.ToLower(s))
	}

	// If the string has no underscores, capitalize it and leave it be.
	if !strings.ContainsRune(s, '_') {
		return strings.Title(s)
	}

	// Now split the name at underscores, capitalize the first
	// letter of each chunk, and smush'em back together.
	chunks := strings.Split(s, "_")
	for i, chunk := range chunks {
		chunks[i] = strings.Title(strings.ToLower(chunk))
	}
	return strings.Join(chunks, "")
}
