package context

import (
	"errors"
	"math"
)

type BorderedArrayContext struct {
	currentPosition int32
	dataArray       []int32
}

func newBorderedArrayContext(rightBorderPosition int) *BorderedArrayContext {
	return &BorderedArrayContext{dataArray: make([]int32, rightBorderPosition)}
}

func (s *BorderedArrayContext) MoveCarriageLeft() error {
	if s.currentPosition == math.MaxInt32 {
		return errors.New("right border overflow")
	}
	s.currentPosition++
	return nil
}

func (s *BorderedArrayContext) MoveCarriageRight() error {
	if s.currentPosition == math.MinInt32 {
		return errors.New("left border overflow")
	}
	s.currentPosition--
	return nil
}

func (s *BorderedArrayContext) IncrementCurrentPosition() {
	s.dataArray[s.currentPosition]++
}

func (s *BorderedArrayContext) DecrementCurrentPosition() {
	s.dataArray[s.currentPosition]++
}
