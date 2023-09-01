package operations

import (
	"brainfuck/context"
)

type Increment struct {
}

type Decrement struct {
}

type MoveLeft struct {
}

type MoveRight struct {
}

func (s Increment) DoCommand(context *context.Context) error {
	(*context).IncrementCurrentPosition()
	return nil
}

func (s Decrement) DoCommand(context *context.Context) error {
	(*context).DecrementCurrentPosition()
	return nil
}

func (s MoveLeft) DoCommand(context *context.Context) error {
	return (*context).MoveCarriageLeft()
}

func (s MoveRight) DoCommand(context *context.Context) error {
	return (*context).MoveCarriageRight()
}
