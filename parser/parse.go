package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"sort"
)

const (
	GoFilesSuffix = ".go"
)

func ReadContent(filename, path string) (string, error) {
	fmt.Println(path)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return "", nil
	}
	fset := token.NewFileSet() // positions are relative to fset
	a, err := parser.ParseFile(fset, "", string(fd), 0)
	if err != nil {
		panic(err)
	}
	var arr FuncArray
	ast.Inspect(a, func(node ast.Node) bool {
		fn, ok := node.(*ast.FuncDecl)
		if ok {
			var receiver string
			if fn.Recv != nil {
				if st, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
					if ident, ok := st.X.(*ast.Ident); ok {
						//fmt.Printf("%s", ident.Name)
						receiver = ident.Name
					}
				}
				//fmt.Printf("%v\n---\n", fn.Recv.List[0].Type)
			}

			//fmt.Printf("%sfunction declaration found on line %d: %s\n", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			fnSt := &Func{
				Receiver: receiver,
				Exported: fn.Name.IsExported(),
				Name:     fn.Name.Name,
				Start:    fset.Position(fn.Pos()).Line,
				Length:   fset.Position(fn.Body.Rbrace).Line - fset.Position(fn.Body.Lbrace).Line,
				Lbrace:   fset.Position(fn.Body.Lbrace).Line,
				Rbrace:   fset.Position(fn.Body.Rbrace).Line,
			}
			arr = append(arr, fnSt)

			return true
		}
		return true
	})
	sort.Sort(arr)
	for _, v := range arr {
		fmt.Printf("%#v\n", v)
	}
	return "", nil
}

type Func struct {
	Receiver string
	Exported bool
	Name     string
	Start    int
	Length   int
	Lbrace   int
	Rbrace   int
}

type FuncArray []*Func

func (f FuncArray) Len() int {
	return len(f)
}

func (f FuncArray) Less(i, j int) bool {
	return f[i].Length > f[j].Length
}

func (f FuncArray) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
