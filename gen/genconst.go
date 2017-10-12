package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"go/ast"
	"text/template"
	"os"
	"path/filepath"
)

/*
Contains names of consts in gowpd.go
 */
type Consts struct {
	HRESULTS []string
	CLSIDS []string
	IIDS []string
	PROPERTYKEYS []string
	GUIDS []string
}

func main() {
	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, "gowpd.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	af.Doc = nil
	af.Imports = nil
	af.Comments = nil

	var consts Consts
	for _, decl := range af.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if gd.Tok != token.CONST {
			continue
		}

		var lastType ast.Expr
		for _, spec := range gd.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			var curType ast.Expr = vs.Type
			if curType == nil {
				curType = lastType
			}

			if curType == nil {
				continue
			}

			typeName := curType.(*ast.Ident).Name
			name := vs.Names[0].Name
			switch typeName {
			case "HRESULT":
				consts.HRESULTS = append(consts.HRESULTS, name)
			case "CLSID":
				consts.CLSIDS = append(consts.CLSIDS, name)
			case "IID":
				consts.IIDS = append(consts.IIDS, name)
			case "PropertyKey":
				consts.PROPERTYKEYS = append(consts.PROPERTYKEYS, name)
			case "GUID":
				consts.GUIDS = append(consts.GUIDS, name)
			}
			log.Println(typeName, name)

			lastType = curType
		}
	}

	templateBytes, err := ioutil.ReadFile("gen/consts.go.template")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("consts.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(filepath.Abs(f.Name()))

	constsTemplate := template.Must(template.New("consts").Parse(string(templateBytes)))
	err = constsTemplate.Execute(f, consts)
	if err != nil {
		log.Fatal(err)
	}
}
