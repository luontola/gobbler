// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gobbler

import (
	"container/vector"
)


type StringSet []string

func (this *StringSet) Add(value string) {
	items := this.asVector()
	for i := 0; i < items.Len(); i++ {
		if items.At(i) == value {
			return
		}
	}
	items.Push(value)
}

func (this *StringSet) asVector() *vector.StringVector {
	return (*vector.StringVector)(this)
}

func (this *StringSet) ToArray() []string {
	return *this
}

