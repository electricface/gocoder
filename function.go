package gocoder

import (
	"fmt"
	"strings"
)

type Function struct {
	Desc string
	Name string

	IsMethod bool
	RecvName string
	RecvType string

	Arg    string
	Return string
	Body   FuncBody
}

type FuncBody []string

func (fb FuncBody) String() string {
	strv := []string(fb)
	return strings.Join(strv, "\n")
}

// new line in body
func (fn *Function) Line(l string) {
	fn.Body = append(fn.Body, l)
}

func (fn *Function) Break() {
	fn.Body = append(fn.Body, "")
}

func (fn *Function) Linef(format string, a ...interface{}) {
	fn.Body = append(fn.Body, fmt.Sprintf(format, a...))
}

func (f *Function) String() string {
	var recv string
	if f.IsMethod {
		recv = "(" + f.RecvName + " " + f.RecvType + ") "
	}
	// return example:
	// error
	// int,error
	// a int
	// a int,err error
	returns := f.Return
	if strings.ContainsAny(returns, " ,") {
		returns = "(" + returns + ")"
	}
	header := "func " + recv + f.Name + "(" + f.Arg + ") " + returns
	return Comment(f.Desc) + header + " {\n" + f.Body.String() + "\n}\n"
}

// args: func name string, recvName string, desc string
func newMethod(typeName string, p bool, args ...string) *Function {
	nargs := len(args)
	if !(1 <= nargs && nargs <= 3) {
		panic("length of args must be 1 to 3")
	}
	recvType := typeName
	if p {
		recvType = "*" + recvType
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

func NewMethod(typeName string, args ...string) *Function {
	return newMethod(typeName, false, args...)
}

func NewMethodP(typeName string, args ...string) *Function {
	return newMethod(typeName, true, args...)
}
