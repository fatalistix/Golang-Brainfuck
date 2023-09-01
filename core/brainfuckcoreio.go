package core

import (
	"brainfuck/context"
	"brainfuck/factory"
	"brainfuck/format"
	"fmt"
	"io"
	"os"
)

type BrainfuckCoreIO struct {
	context        context.Context
	factory        factory.Factory
	formatter      format.Formatter
	in             io.Reader
	out            io.Writer
	singleIO       bool
	breakOnUnknown bool
}

func MakeBrainfuckCore(context context.Context, factory factory.Factory, formatter format.Formatter, in io.Reader, out io.Writer) BrainfuckCoreIO {
	return BrainfuckCoreIO{context, factory, formatter, in, out}
}

func MakeBrainfuckCoreWithStds(context context.Context, factory factory.Factory, formatter format.Formatter) BrainfuckCoreIO {
	return BrainfuckCoreIO{context, factory, formatter, os.Stdin, os.Stdout}
}

func (bfc BrainfuckCoreIO) StartInterpreting() error {
	var commands string
	var err error
	for {
		_, err = fmt.Fprint(bfc.out, bfc.formatter.Format(format.Read, ""))
		if err != nil {
			return err
		}
		_, err = fmt.Fscanln(bfc.in, &commands)
		if err != nil {
			return err
		}

		for _, c := range commands {
			operation, err := bfc.factory.CreateOperationInstance(string(c))
			if err != nil {
				fmt.Fprintln(bfc.out, bfc.formatter.Format(format.Status()))
			}
			operation.DoCommand(bfc.context)
		}
	}
}
