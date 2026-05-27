package bytes

import (
	"regexp"
)

type (
	// Bytes struct
	Bytes struct{}
)

// binary units (IEC 60027)
const (
	_ = 1.0 << (10 * iota) // ignore first value by assigning to blank identifier
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

// decimal units (SI international system of units)
const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
)

var (
	patternBinary  = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)\s?([KMGTPE]iB?)$`)
	patternDecimal = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)\s?([KMGTPE]B?|B?)$`)
	global         = New()
)

// New creates a Bytes instance.
func New() *Bytes {
	_ = "STUB: not implemented"

	// Format formats bytes integer to human readable string according to IEC 60027.
	// For example, 31323 bytes will return 30.59KB.
	return nil
}

func (b *Bytes) Format(value int64) string { _ = "STUB: not implemented"; return "" }

// FormatBinary formats bytes integer to human readable string according to IEC 60027.
// For example, 31323 bytes will return 30.59KB.
func (*Bytes) FormatBinary(value int64) string { _ = "STUB: not implemented"; return "" }

// FormatDecimal formats bytes integer to human readable string according to SI international system of units.
// For example, 31323 bytes will return 31.32KB.
func (*Bytes) FormatDecimal(value int64) string { _ = "STUB: not implemented"; return "" }

// Parse parses human readable bytes string to bytes integer.
// For example, 6GiB (6Gi is also valid) will return 6442450944, and
// 6GB (6G is also valid) will return 6000000000.
func (b *Bytes) Parse(value string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ParseBinary parses human readable bytes string to bytes integer.
// For example, 6GiB (6Gi is also valid) will return 6442450944.
func (*Bytes) ParseBinary(value string) (i int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ParseDecimal parses human readable bytes string to bytes integer.
// For example, 6GB (6G is also valid) will return 6000000000.
func (*Bytes) ParseDecimal(value string) (i int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Format wraps global Bytes's Format function.
func Format(value int64) string { _ = "STUB: not implemented"; return "" }

// FormatBinary wraps global Bytes's FormatBinary function.
func FormatBinary(value int64) string { _ = "STUB: not implemented"; return "" }

// FormatDecimal wraps global Bytes's FormatDecimal function.
func FormatDecimal(value int64) string { _ = "STUB: not implemented"; return "" }

// Parse wraps global Bytes's Parse function.
func Parse(value string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
