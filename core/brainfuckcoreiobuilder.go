package core

import (
	"io"
)

type BrainfuckCoreIOBuilder interface {
	Build() BrainfuckCoreIO
	SetSingleIO(singleIO bool) BrainfuckCoreIOBuilder
	SetBreakOnUnknown(breakOnUnknown bool) BrainfuckCoreIOBuilder
	SetCache(cache bool) BrainfuckCoreIOBuilder
	SetInput(input io.Reader) BrainfuckCoreIOBuilder
	SetOutput(output io.Writer) BrainfuckCoreIOBuilder
}
