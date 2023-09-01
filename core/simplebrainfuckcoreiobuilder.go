package core

import (
	"brainfuck/context"
	"brainfuck/factory"
	"brainfuck/format"
	"io"
)

type SimpleBrainfuckCoreIOBuilder struct {
	singleIO       bool
	breakOnUnknown bool
	cache bool
	context        context.Context
	factory        factory.Factory
	formatter      format.Formatter
	input          io.Reader
	output         io.Writer
}

func newSimpleBrainfuckCoreIOBuilder() *SimpleBrainfuckCoreIOBuilder {
	return &SimpleBrainfuckCoreIOBuilder{false, false, false, }
}

func (s *SimpleBrainfuckCoreIOBuilder)
