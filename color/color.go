package color

import (
	"io"
)

type (
	inner func(interface{}, []string, *Color) string
)

// Color styles
const (
	// Blk Black text style
	Blk = "30"
	// Rd red text style
	Rd = "31"
	// Grn green text style
	Grn = "32"
	// Yel yellow text style
	Yel = "33"
	// Blu blue text style
	Blu = "34"
	// Mgn magenta text style
	Mgn = "35"
	// Cyn cyan text style
	Cyn = "36"
	// Wht white text style
	Wht = "37"
	// Gry grey text style
	Gry = "90"

	// BlkBg black background style
	BlkBg = "40"
	// RdBg red background style
	RdBg = "41"
	// GrnBg green background style
	GrnBg = "42"
	// YelBg yellow background style
	YelBg = "43"
	// BluBg blue background style
	BluBg = "44"
	// MgnBg magenta background style
	MgnBg = "45"
	// CynBg cyan background style
	CynBg = "46"
	// WhtBg white background style
	WhtBg = "47"

	// R reset emphasis style
	R = "0"
	// B bold emphasis style
	B = "1"
	// D dim emphasis style
	D = "2"
	// I italic emphasis style
	I = "3"
	// U underline emphasis style
	U = "4"
	// In inverse emphasis style
	In = "7"
	// H hidden emphasis style
	H = "8"
	// S strikeout emphasis style
	S = "9"
)

var (
	black   = outer(Blk)
	red     = outer(Rd)
	green   = outer(Grn)
	yellow  = outer(Yel)
	blue    = outer(Blu)
	magenta = outer(Mgn)
	cyan    = outer(Cyn)
	white   = outer(Wht)
	grey    = outer(Gry)

	blackBg   = outer(BlkBg)
	redBg     = outer(RdBg)
	greenBg   = outer(GrnBg)
	yellowBg  = outer(YelBg)
	blueBg    = outer(BluBg)
	magentaBg = outer(MgnBg)
	cyanBg    = outer(CynBg)
	whiteBg   = outer(WhtBg)

	reset     = outer(R)
	bold      = outer(B)
	dim       = outer(D)
	italic    = outer(I)
	underline = outer(U)
	inverse   = outer(In)
	hidden    = outer(H)
	strikeout = outer(S)

	global = New()
)

func outer(n string) inner { _ = "STUB: not implemented"; return *new(inner) }

// TODO: Drop fmt to boost performance?

type (
	Color struct {
		output   io.Writer
		disabled bool
	}
)

// New creates a Color instance.
func New() (c *Color) { _ = "STUB: not implemented"; return nil }

// Output returns the output.
func (c *Color) Output() io.Writer {
	_ = "STUB: not implemented"

	// SetOutput sets the output.
	return *new(io.Writer)
}

func (c *Color) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

// Disable disables the colors and styles.
func (c *Color) Disable() {
	_ = "STUB: not implemented"

	// Enable enables the colors and styles.
	return
}

func (c *Color) Enable() {
	_ = "STUB: not implemented"

	// Print is analogous to `fmt.Print` with termial detection.
	return
}

func (c *Color) Print(args ...interface{}) { _ = "STUB: not implemented"; return }

// Println is analogous to `fmt.Println` with termial detection.
func (c *Color) Println(args ...interface{}) { _ = "STUB: not implemented"; return }

// Printf is analogous to `fmt.Printf` with termial detection.
func (c *Color) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (c *Color) Black(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Red(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func (c *Color) Green(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Yellow(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Blue(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Magenta(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Cyan(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) White(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Grey(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) BlackBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) RedBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) GreenBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) YellowBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) BlueBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) MagentaBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) CyanBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) WhiteBg(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Reset(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Bold(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Dim(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func (c *Color) Italic(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Underline(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Inverse(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Hidden(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Color) Strikeout(msg interface{}, styles ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// Output returns the output.
func Output() io.Writer {
	_ = "STUB: not implemented"
	return *

	// SetOutput sets the output.
	new(io.Writer)
}

func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func Disable() { _ = "STUB: not implemented"; return }

func Enable() {
	_ = "STUB: not implemented"

	// Print is analogous to `fmt.Print` with termial detection.
	return
}

func Print(args ...interface{}) { _ = "STUB: not implemented"; return }

// Println is analogous to `fmt.Println` with termial detection.
func Println(args ...interface{}) { _ = "STUB: not implemented"; return }

// Printf is analogous to `fmt.Printf` with termial detection.
func Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Black(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Red(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Green(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Yellow(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Blue(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Magenta(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Cyan(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func White(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Grey(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func BlackBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func RedBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func GreenBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func YellowBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func BlueBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func MagentaBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func CyanBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func WhiteBg(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Reset(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Bold(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Dim(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Italic(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Underline(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Inverse(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Hidden(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }

func Strikeout(msg interface{}, styles ...string) string { _ = "STUB: not implemented"; return "" }
