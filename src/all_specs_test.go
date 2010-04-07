// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	"gospec"
	"testing"
)


func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(FileDependenciesSpec)
	r.AddSpec(PackageDependenciesSpec)
	r.AddSpec(StringSetSpec)
	r.AddSpec(TemporaryDirectorySpec)
	gospec.MainGoTest(r, t)
}

