package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// GetStructList 获取包下面所有结构体  pkgName 要遍历的包名  pkgPath 包的相对路径
func GetStructList(pkgName string, pkgPath string) {
	fmt.Println("........test")
	//pkgName := "model" // 要遍历的包名
	// 解析包中所有源代码文件
	fset := token.NewFileSet()
	//pkgs, err := parser.ParseDir(fset, "./model", nil, parser.AllErrors)
	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// 遍历指定包中所有类型定义
	for _, pkg := range pkgs {
		if pkg.Name != pkgName {
			continue
		}

		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if typ, ok := decl.(*ast.GenDecl); ok && typ.Tok == token.TYPE {
					for _, spec := range typ.Specs {
						if st, ok := spec.(*ast.TypeSpec); ok && st.Type != nil {
							if _, ok := st.Type.(*ast.StructType); ok {
								//todo 如何处理数据
								fmt.Printf("Found struct type %q in file %q\n", st.Name.Name, fset.Position(typ.Pos()).Filename)
							}
						}
					}
				}
			}
		}
	}
}
