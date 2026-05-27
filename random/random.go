package random

import (
	"sync"
)

type (
	Random struct {
		readerPool sync.Pool
	}
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
)

var (
	global = New()
)

func New() *Random {
	_ = "STUB: not implemented"
	// https://tip.golang.org/doc/go1.19#:~:text=Read%20no%20longer%20buffers%20random%20data%20obtained%20from%20the%20operating%20system%20between%20calls
	// sync.Pool must not be copied after first use; construct it directly
	// on the struct to avoid copying from a local var.
	return nil
}

func (r *Random) String(length uint8, charsets ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// perf: avoid read from rand.Reader many times

// security note:
// we can't just simply do b[i]=charset[rb%byte(charsetLen)],
// for example, when charsetLen is 52, and rb is [0, 255], 256 = 52 * 4 + 48.
// this will make the first 48 characters more possibly to be generated then others.
// so we have to skip bytes when rb > maxByte

// Skip this number to avoid bias.

func String(length uint8, charsets ...string) string { _ = "STUB: not implemented"; return "" }
