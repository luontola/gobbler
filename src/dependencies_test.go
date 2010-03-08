// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	"gospec"
	. "gospec"
)


func FileDependenciesSpec(c gospec.Context) {

	c.Specify("Dependencies area read from a file's imports:", func() {
	
		c.Specify("no imports", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
			`)
			c.Expect(deps, ContainsExactly, Values())
		})
		c.Specify("one import", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import "dep1"
			`)
			c.Expect(deps, ContainsExactly, Values("dep1"))
		})
		c.Specify("many imports", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import "dep1"
				import "dep2"
			`)
			c.Expect(deps, ContainsExactly, Values("dep1", "dep2"))
		})
		c.Specify("import blocks", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import (
					"dep1"
					"dep2"
				)
			`)
			c.Expect(deps, ContainsExactly, Values("dep1", "dep2"))
		})
		c.Specify("named import", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import name "dep1"
			`)
			c.Expect(deps, ContainsExactly, Values("dep1"))
		})
		c.Specify(". import", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import . "dep1"
			`)
			c.Expect(deps, ContainsExactly, Values("dep1"))
		})
		c.Specify("_ import", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import _ "dep1"
			`)
			c.Expect(deps, ContainsExactly, Values("dep1"))
		})
	})
	
	c.Specify("Duplicate imports are reported only once", func() {
		deps, _ := GetFileDependencies("file.go", `
			package mypkg
			import a "dep1"
			import b "dep1"
		`)
		c.Expect(deps, ContainsExactly, Values("dep1"))
	})
	
	c.Specify("Syntax errors are handled gracefully:", func() {
	
		c.Specify("missing quotes", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import dep1
			`)
			c.Expect(deps, ContainsExactly, Values())
		})
		c.Specify("garbage after the import", func() {
			deps, _ := GetFileDependencies("file.go", `
				package mypkg
				import "dep1" garbage
			`)
			c.Expect(deps, ContainsExactly, Values())
		})
		c.Specify("empty file", func() {
			deps, _ := GetFileDependencies("file.go", `
			`)
			c.Expect(deps, ContainsExactly, Values())
		})
	})
}

func PackageDependenciesSpec(c Context) {
	/*
	pkg, _ := CreateTempDir()
	defer pkg.Dispose()
	
	c.Specify("Given a package with no files", func() {
		dependencies := GetPackageDependencies(pkg.Path())
		
		c.Specify("it has no dependencies", func() {
			c.Expect(dependencies, ContainsExactly, Values())
		})
	})
	
	c.Specify("Given a package with one .go file", func() {
		pkg.WriteFile("file.go", `
			package main
			import "dependency"
		`)
		dependencies := GetPackageDependencies(pkg.Path())
		
		c.Specify("it has the dependencies of that one file", func() {
			c.Expect(dependencies, ContainsExactly, Values("dependency"))
		})
	})
	
	c.Specify("Given a package with many .go files", func() {
		pkg.WriteFile("file1.go", `
			package main
			import "dependency1"
			import "common/dependency"
		`)
		pkg.WriteFile("file2.go", `
			package main
			import "dependency2"
			import "common/dependency"
		`)
		dependencies := GetPackageDependencies(pkg.Path())
		
		c.Specify("it has the combined dependencies of all those files", func() {
			c.Expect(dependencies, ContainsExactly, Values("dependency1", "dependency2", "common/dependency"))
		})
	})
	*/
}

