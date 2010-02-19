// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	. "gospec"
)


func StringSetSpec(c Context) {
	set := new(StringSet)

	c.Specify("An empty set contains nothing", func() {
		c.Expect(set.ToArray(), ContainsExactly, Values())
	})
	c.Specify("Distinct values can be added to a set", func() {
		set.Add("a")
		set.Add("b")
		c.Expect(set.ToArray(), ContainsExactly, Values("a", "b"))
	})
	c.Specify("Duplicate values cannot be added to a set", func() {
		set.Add("a")
		set.Add("a")
		c.Expect(set.ToArray(), ContainsExactly, Values("a"))
	})
}

