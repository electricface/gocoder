package main

import (
	"github.com/electricface/gocoder"
	"os"
)

func main() {
	sourceFile := gocoder.NewSourceFile()
	sourceFile.Package.Name = "main"
	sourceFile.Import("fmt")
	sourceFile.Import("fmt")
	sourceFile.Import("fmt")
	sourceFile.ImportAs("io", "io0")
	sourceFile.ImportAs("abc/def", ".")
	sourceFile.ImportAs("abc/def", "_")

	typeAlias0 := &gocoder.TypeAlias{
		Desc: "a type alias define",
		New:  "Strv",
		Old:  "[]string",
	}
	sourceFile.Add(typeAlias0)

	struct0 := &gocoder.Struct{
		Desc: "a struct define",
		Name: "StructA",
	}
	struct0.AddField("a", "int", "name a")
	struct0.AddField("B", "string")
	sourceFile.Add(struct0)

	sfn1 := struct0.NewMethodP("GetA", "", "test new method p")
	sfn1.Return = "int"
	sfn1.Line("return self.a")
	sourceFile.Add(sfn1)

	sfn2 := struct0.NewMethodV("Type", "_")
	sfn2.Return = "string"
	sfn2.Line(`return "const"`)
	sourceFile.Add(sfn2)

	const0 := &gocoder.Const{
		Desc:  "a const define",
		Name:  "INT_MAX",
		Value: 1,
	}

	const1 := &gocoder.Const{
		Name:  "STR_NAME",
		Value: "swt",
	}

	const2 := &gocoder.Const{
		Name:  "FLOAT_A",
		Typ:   "float64",
		Value: 1.2345,
	}

	constGroup := &gocoder.DefineGroup{
		Desc: "const block",
		Type: "const",
		Items: []gocoder.DefineItem{
			const0, const1, const2,
		},
	}
	sourceFile.Add(constGroup)

	func1 := &gocoder.Function{
		Desc:   "a named function define",
		Name:   "TestA",
		Arg:    "a int, b int",
		Return: "int,int",
		//Code: `return a+b`,
	}
	func1.Line("a = a + b + 1")
	func1.Linef("b += %d", 123)
	func1.Break()
	func1.Line("return a+b, b-a")
	sourceFile.Add(func1)
	sourceFile.Add(func1)

	var0 := &gocoder.Var{
		Name:  "logger",
		Value: `log.NewLogger("xxx")`,
	}
	sourceFile.Add(var0)

	var1 := &gocoder.Var{
		Name: "state",
		Typ:  "int",
	}
	sourceFile.Add(var1)

	sourceFile.Add(&gocoder.DefineGroup{
		Type: "var",
		Items: []gocoder.DefineItem{
			var0, var1,
		},
	})

	sourceFile.Line("// a comment line")
	sourceFile.Line("// a comment line")
	sourceFile.Linef("// my name is %s", "swt")

	constgroup1 := sourceFile.AddConstGroup("test AddConstGroup")
	constgroup1.Add(&gocoder.Const{
		Name:  "abcdef",
		Value: 1,
	})
	constgroup1.Add(&gocoder.Const{
		Name:  "qwert",
		Value: 2,
	})

	sourceFile.Add(&gocoder.Const{
		Name:  "output",
		Value: true,
	})

	sourceFile.WriteTo(os.Stdout)
	os.Stdout.Close()
}
