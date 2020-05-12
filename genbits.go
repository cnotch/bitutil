// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const templateText = `// Code generated using genbits.go; DO NOT EDIT.

// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bitutil
{{range .}}
// Test{{.Len}} returns the ith bit value of an uint{{.Len}}.
func Test{{.Len}}(v uint{{.Len}}, i int) bool {
	if i < 0 || i > {{.MaxIndex}} {
		return false
	}
	return v&(1<<({{.MaxIndex}}-i)) != 0
}

// Set{{.Len}} set the ith bit value of an uint{{.Len}} to 1, and returns the new uint{{.Len}}.
func Set{{.Len}}(v uint{{.Len}}, i int) uint{{.Len}} {
	if i < 0 || i > {{.MaxIndex}} {
		return v
	}
	return v | (1 << ({{.MaxIndex}} - i))
}

// Clear{{.Len}} clear the ith bit value of an uint{{.Len}} to 0, and returns the new uint{{.Len}}.
func Clear{{.Len}}(v uint{{.Len}}, i int) uint{{.Len}} {
	if i < 0 || i > {{.MaxIndex}} {
		return v
	}
	return v & ^(1 << ({{.MaxIndex}} - i))
}

// SetTo{{.Len}} set the ith bit value of an uint{{.Len}}, and returns the new uint{{.Len}}.
func SetTo{{.Len}}(v uint{{.Len}}, i int, b bool) uint{{.Len}} {
	if b {
		return Set{{.Len}}(v, i)
	}
	return Clear{{.Len}}(v, i)
}

// Flip{{.Len}} flip the ith bit value of an uint{{.Len}}, and returns the new uint{{.Len}}.
func Flip{{.Len}}(v uint{{.Len}}, i int) uint{{.Len}} {
	if i < 0 || i > {{.MaxIndex}} {
		return v
	}
	return v ^(1 << ({{.MaxIndex}} - i))
}

// Sub{{.Len}} returns the copy uint{{.Len}} of v[begin:end].
func Sub{{.Len}}(v uint{{.Len}}, begin, end int) uint{{.Len}} {
	if begin < 0 || begin > {{.Len}} || end < 0 || end > {{.Len}} {
		return 0
	}
	return (v << begin) >> ({{.Len}} - end + begin)
}

// Left{{.Len}} returns the copy uint{{.Len}} of v[:w].
func Left{{.Len}}(v uint{{.Len}}, w int) uint{{.Len}} {
	if w < 0 || w > {{.Len}} {
		return 0
	}
	return v >> ({{.Len}} - w)
}

// Right{{.Len}} returns the copy uint{{.Len}} of v[{{.Len}}-w:].
func Right{{.Len}}(v uint{{.Len}}, w int) uint{{.Len}} {
	if w < 0 || w > {{.Len}} {
		return 0
	}
	return (v << ({{.Len}} - w)) >> ({{.Len}} - w)
}

// Replace{{.Len}} returns a copy of the v with v[begin:end] instances of old replaced by new[{{.Len}}-end+begin:].
func Replace{{.Len}}(v uint{{.Len}}, begin, end int, new uint{{.Len}}) uint{{.Len}} {
	if begin < 0 || begin > {{.Len}} || end < 0 || end > {{.Len}} {
		return v
	}
	mask := (uint{{.Len}}({{.Max}}) << ({{.Len}} - end + begin)) >> begin
	return (v & ^mask) | ((new << ({{.Len}} - end)) & mask)
}
{{end}}`
	type BitsType struct {
		Len      int
		MaxIndex int
		Max      string
	}
	var typs = []BitsType{
		{8, 7, "0xff"},
		{16, 15, "0xffff"},
		{32, 31, "0xffffffff"},
		{64, 63, "0xffffffffffffffff"},
	}

	t := template.Must(template.New("bits").Parse(templateText))
	f, err := os.OpenFile("uint.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("executing template:", err)
	}

	defer f.Close()

	// Execute the template for each bitstype.
	err = t.Execute(f, typs)
	if err != nil {
		log.Println("executing template:", err)
	}
}
