// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	"go/ast"
	"go/parser"
	"os"
	"regexp"
)


func GetFileDependencies(filename string, src interface{}) ([]string, os.Error) {
	file, err := parser.ParseFile(filename, src, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}
	
	visitor := NewDependencyVisitor()
	ast.Walk(visitor, file)
	return visitor.GetImports(), nil
}


func NewDependencyVisitor() *DependencyVisitor {
	return &DependencyVisitor{
		new(StringSet),
	}
}

type DependencyVisitor struct {
	imports *StringSet
}

func (this *DependencyVisitor) GetImports() []string {
	return *this.imports
}

func (this *DependencyVisitor) Visit(node interface{}) ast.Visitor {
	switch n := node.(type) {
	case *ast.ImportSpec:
		this.addImports(n)
	}
	return this
}

func (this *DependencyVisitor) addImports(importSpec *ast.ImportSpec) {
	for _, path := range importSpec.Path {
		this.imports.Add(stripQuotes(string(path.Value)))
	}
}

var unquote = regexp.MustCompile(`"(.+)"`)

func stripQuotes(quoted string) string {
	matches := unquote.MatchStrings(quoted)
	return matches[1]
}


func GetPackageDependencies(packagePath string) []string {
	return nil
}

