package factory

import (
	"brainfuck/context"
)

type Operation interface {
	DoCommand(context *context.Context) error
}

type Factory interface {
	CreateOperationInstance(command string) (Operation, error)
	Register(command string, operation Operation) error
}
