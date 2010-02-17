// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	. "gospec"
)


func PackageDependenciesSpec(c Context) {
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
}

