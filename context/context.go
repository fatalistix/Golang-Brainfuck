package context

type Context interface {
	MoveCarriageLeft() error
	MoveCarriageRight() error

	IncrementCurrentPosition()
	DecrementCurrentPosition()
}
