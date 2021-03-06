// Copyright 2015 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"strings"
	"unicode"

	. "github.com/limetext/backend"
)

type (
	ToggleComment struct {
		DefaultCommand
	}
)

func (c *ToggleComment) Run(v *View, e *Edit) error {
	// TODO: Comment the line if we only have a cursor.
	// TODO: Expand the selection after altering it.
	// TODO: Align the comment characters for multiline selections.
	// TODO: Get the comment value from the Textmate files.
	comm := "//"

	for _, r := range v.Sel().Regions() {
		if r.Size() != 0 {
			t := v.Substr(r)

			trim := strings.TrimLeftFunc(t, unicode.IsSpace)
			if strings.HasPrefix(trim, comm) {
				repl := comm
				if strings.HasPrefix(trim, comm+" ") {
					repl += " "
				}

				t = strings.Replace(t, repl, "", 1)
			} else {
				t = strings.Replace(t, trim, comm+" "+trim, 1)
			}

			v.Replace(e, r, t)
		}
	}

	return nil
}

func init() {
	register([]Command{
		&ToggleComment{},
	})
}
