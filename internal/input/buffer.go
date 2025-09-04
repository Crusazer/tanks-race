package input

import (
	"time"

	"github.com/Crusazer/tanks-race/pkg/math"
)

type InputFrame[T comparable] struct {
    FrameNumber int
    Actions     map[T]bool
    MousePosition math.Vector2
    Timestamp   time.Time
}

type InputBuffer[T comparable] struct {
    frames []InputFrame[T]
    maxSize int
}

func NewInputBuffer[T comparable] (maxSize int) *InputBuffer[T] {
    return &InputBuffer[T]{
        frames:  make([]InputFrame[T], 0, maxSize),
        maxSize: maxSize,
    }
}

func (ib *InputBuffer[T]) AddFrame(frame InputFrame[T]) {
    ib.frames = append(ib.frames, frame)
    
    // Ограничиваем размер буфера
    if len(ib.frames) > ib.maxSize {
        ib.frames = ib.frames[1:]
    }
}

func (ib *InputBuffer[T]) GetFrame(frameNumber int) *InputFrame[T] {
    for i := len(ib.frames) - 1; i >= 0; i-- {
        if ib.frames[i].FrameNumber == frameNumber {
            return &ib.frames[i]
        }
    }
    return nil
}

func (ib *InputBuffer[T]) Clear() {
    ib.frames = ib.frames[:0]
}
