package gocoder

import (
	"strings"
)

type Struct struct {
	Desc   string
	Name   string
	Fields []*StructField
}

func (s *Struct) String() string {
	var parts []string
	for _, item := range s.Fields {
		parts = append(parts, item.String())
	}
	return Comment(s.Desc) + "type " + s.Name + " struct {\n" +
		strings.Join(parts, "\n") + "\n}\n"
}

func (s *Struct) AddField(args ...string) {
	field := new(StructField)
	nargs := len(args)
	if !(nargs == 2 || nargs == 3) {
		panic("length of args must be 2 or 3")
	}
	field.Name = args[0]
	field.Type = args[1]
	if nargs == 3 {
		field.Desc = args[2]
	}

	s.Fields = append(s.Fields, field)
}

func (s *Struct) AppendField(field *StructField) {
	s.Fields = append(s.Fields, field)
}

// default recv name is "self"
// name string, recvName string, desc string
func (s *Struct) newMethod(p bool, args ...string) *Function {
	nargs := len(args)
	if !(1 <= nargs && nargs <= 3) {
		panic("length of args must be 1 to 3")
	}
	recvType := s.Name
	if p {
		recvType = "*" + s.Name
	}
	fn := &Function{
		Name:     args[0],
		IsMethod: true,
		RecvType: recvType,
		RecvName: "self",
	}
	// narg is 2 or 3
	if nargs > 1 && args[1] != "" {
		fn.RecvName = args[1]
	}

	// narg is 3
	if nargs > 2 {
		fn.Desc = args[2]
	}
	return fn
}

func (s *Struct) NewMethodP(args ...string) *Function {
	return s.newMethod(true, args...)
}

func (s *Struct) NewMethodV(args ...string) *Function {
	return s.newMethod(false, args...)
}

type StructField struct {
	Desc string
	Name string
	Type string
}

func (sf *StructField) String() string {
	return Comment(sf.Desc) + sf.Name + " " + sf.Type
}
