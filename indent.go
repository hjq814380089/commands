// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"strings"

	. "github.com/limetext/backend"
	. "github.com/limetext/text"
)

type (
	// The IndentCommand increments indentation of selection.
	Indent struct {
		DefaultCommand
	}

	// The UnindentCommand decrements indentation of selection.
	Unindent struct {
		DefaultCommand
	}
)

func (c *Indent) Run(v *View, e *Edit) error {
	indent := "\t"
	if t := v.Settings().Bool("translate_tabs_to_spaces", false); t {
		indent = strings.Repeat(" ", v.Settings().Int("tab_size", 4))
	}
	sel := v.Sel()

	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		start_row, _ := v.RowCol(r.Begin())
		end_row, _ := v.RowCol(r.End())
		for row := start_row; row <= end_row; row++ {
			// Insert an indent at the beginning of the line
			pos := v.TextPoint(row, 0)
			v.Insert(e, pos, indent)
		}
	}
	return nil
}

func (c *Unindent) Run(v *View, e *Edit) error {
	tab_size := v.Settings().Int("tab_size", 4)
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		start_row, _ := v.RowCol(r.Begin())
		end_row, _ := v.RowCol(r.End())
		for row := start_row; row <= end_row; row++ {
			pos := v.TextPoint(row, 0)
			// Get the first at the beginning of the line (as many as defined by tab_size)
			sub := v.Substr(Region{pos, pos + tab_size})
			if len(sub) == 0 {
				continue
			}
			to_remove := 0
			if sub[0] == byte('\t') {
				// Case 1: the first character is a tab, remove only it
				to_remove = 1
			} else if sub[0] == byte(' ') {
				// Case 2: the first character is a space, we remove as much spaces as we can
				to_remove = 1
				for to_remove < len(sub) && sub[to_remove] == byte(' ') {
					to_remove++
				}
			}
			if to_remove > 0 {
				v.Erase(e, Region{pos, pos + to_remove})
			}
		}
	}
	return nil
}

func init() {
	register([]Command{
		&Indent{},
		&Unindent{},
	})
}
