//go:build appengine
// +build appengine

package log

import (
	"io"
)

func output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }
