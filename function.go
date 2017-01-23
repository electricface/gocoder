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
