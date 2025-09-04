package core

type Layout interface {
	Add(w Widget)
	ComputeBounds(parentWidth, parentHeight int)
}

type VerticalFlowLayout struct {
	widgets []Widget
	padding int
	spacing int
}

func NewVerticalFlowLayout(padding, spacing int) *VerticalFlowLayout {
	return &VerticalFlowLayout{padding: padding, spacing: spacing}
}

func (l *VerticalFlowLayout) Widgets() []Widget {
	return l.widgets
}

func (l *VerticalFlowLayout) Add(w Widget) {
	l.widgets = append(l.widgets, w)
}

func (l *VerticalFlowLayout) ComputeBounds(parentWidth, parentHeight int) {
	totalH := l.padding * 2
	for _, w := range l.widgets {
		_, _, _, h := w.Bounds()
		totalH += h
	}
	totalH += l.spacing * (len(l.widgets) - 1)

	y := (parentHeight - totalH) / 2
	for _, w := range l.widgets {
		_, _, wW, wH := w.Bounds()
		x := (parentWidth - wW) / 2
		w.SetBounds(x, y, wW, wH)
		y += wH + l.spacing
	}
}
