package input

import (
	"github.com/Crusazer/tanks-race/pkg/math"
	"time"
)

type InputHistory[T comparable] struct {
	buffer *InputBuffer[T]
}

func NewInputHistory[T comparable](bufferSize int) *InputHistory[T] {
	return &InputHistory[T]{
		buffer: NewInputBuffer[T](bufferSize),
	}
}

// RecordInput теперь принимает map[T]bool вместо InputSystem
func (ih *InputHistory[T]) RecordInput(
	frameNumber int,
	actions map[T]bool,
	mousePosition math.Vector2,
) {
	frame := InputFrame[T]{
		FrameNumber:   frameNumber,
		Actions:       actions,
		MousePosition: mousePosition,
		Timestamp:     time.Now(),
	}

	ih.buffer.AddFrame(frame)
}

func (ih *InputHistory[T]) GetFrame(frameNumber int) *InputFrame[T] {
	return ih.buffer.GetFrame(frameNumber)
}

func (ih *InputHistory[T]) Clear() {
	ih.buffer.Clear()
}

func (ih *InputHistory[T]) RewindToFrame(frameNumber int) {
	// Удаляем все фреймы после указанного
	for i := len(ih.buffer.frames) - 1; i >= 0; i-- {
		if ih.buffer.frames[i].FrameNumber > frameNumber {
			ih.buffer.frames = ih.buffer.frames[:i]
		}
	}
}
