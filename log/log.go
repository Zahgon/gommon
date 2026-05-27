package log

import (
	"io"
	"sync"

	"github.com/valyala/fasttemplate"

	"github.com/labstack/gommon/color"
)

type (
	Logger struct {
		prefix     string
		level      uint32
		skip       int
		output     io.Writer
		template   *fasttemplate.Template
		levels     []string
		color      *color.Color
		bufferPool sync.Pool
		mutex      sync.Mutex
	}

	Lvl uint8

	JSON map[string]interface{}
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	OFF
	panicLevel
	fatalLevel
)

var (
	global        = New("-")
	defaultHeader = `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${short_file}","line":"${line}"}`
)

func init() {
	global.skip = 3
}

func New(prefix string) (l *Logger) { _ = "STUB: not implemented"; return nil }

func (l *Logger) initLevels() { _ = "STUB: not implemented"; return }

func (l *Logger) newTemplate(format string) *fasttemplate.Template {
	_ = "STUB: not implemented"
	return nil
}

func (l *Logger) DisableColor() { _ = "STUB: not implemented"; return }

func (l *Logger) EnableColor() { _ = "STUB: not implemented"; return }

func (l *Logger) Prefix() string { _ = "STUB: not implemented"; return "" }

func (l *Logger) SetPrefix(p string) { _ = "STUB: not implemented"; return }

func (l *Logger) Level() Lvl { _ = "STUB: not implemented"; return *new(Lvl) }

func (l *Logger) SetLevel(level Lvl) { _ = "STUB: not implemented"; return }

func (l *Logger) Output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (l *Logger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func (l *Logger) Color() *color.Color { _ = "STUB: not implemented"; return nil }

func (l *Logger) SetHeader(h string) { _ = "STUB: not implemented"; return }

func (l *Logger) Print(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Printj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Debug(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Debugf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Debugj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Info(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Infoj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Warn(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Warnf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Warnj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Error(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Errorj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatal(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatalj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) Panic(i ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Panicf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Panicj(j JSON) { _ = "STUB: not implemented"; return }

func DisableColor() { _ = "STUB: not implemented"; return }

func EnableColor() { _ = "STUB: not implemented"; return }

func Prefix() string { _ = "STUB: not implemented"; return "" }

func SetPrefix(p string) { _ = "STUB: not implemented"; return }

func Level() Lvl { _ = "STUB: not implemented"; return *new(Lvl) }

func SetLevel(level Lvl) { _ = "STUB: not implemented"; return }

func Output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func SetHeader(h string) { _ = "STUB: not implemented"; return }

func Print(i ...interface{}) { _ = "STUB: not implemented"; return }

func Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Printj(j JSON) { _ = "STUB: not implemented"; return }

func Debug(i ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Debugj(j JSON) { _ = "STUB: not implemented"; return }

func Info(i ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Infoj(j JSON) { _ = "STUB: not implemented"; return }

func Warn(i ...interface{}) { _ = "STUB: not implemented"; return }

func Warnf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Warnj(j JSON) { _ = "STUB: not implemented"; return }

func Error(i ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Errorj(j JSON) { _ = "STUB: not implemented"; return }

func Fatal(i ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalj(j JSON) { _ = "STUB: not implemented"; return }

func Panic(i ...interface{}) { _ = "STUB: not implemented"; return }

func Panicf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Panicj(j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) logJSON(level Lvl, j JSON) { _ = "STUB: not implemented"; return }

func (l *Logger) log(level Lvl, message string, jsonBody bool) { _ = "STUB: not implemented"; return }

// JSON callers route through an extra logJSON wrapper; account for
// that frame so runtime.Caller still lands on the user's code. Keep
// this in sync with logJSON — if the wrapper is ever inlined away
// or moved, drop the increment.

// JSON header

// Text header
