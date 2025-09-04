package core

type Layout interface {
	Add(w Widget)
	Remove(w Widget) // Для динамических меню
	ComputeBounds(parentWidth, parentHeight int)
	Widgets() []Widget
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

func (l *VerticalFlowLayout) Remove(widgetToRemove Widget) {
	for i, w := range l.widgets {
		if w == widgetToRemove { // Сравниваем по указателю
			l.widgets = append(l.widgets[:i], l.widgets[i+1:]...)
			return
		}
	}
}

func (l *VerticalFlowLayout) ComputeBounds(parentWidth, parentHeight int) {
	// Сначала собираем предпочтительные размеры видимых виджетов
	totalPreferredHeight := l.padding * 2
	maxPreferredWidth := 0

	visibleWidgets := []Widget{}
	for _, w := range l.widgets {
		if w.IsVisible() {
			visibleWidgets = append(visibleWidgets, w)
			pW, pH := w.PreferredSize()
			totalPreferredHeight += pH
			if pW > maxPreferredWidth {
				maxPreferredWidth = pW
			}
		}
	}
	totalPreferredHeight += l.spacing * (len(visibleWidgets) - 1)

	// Если есть виджеты, выравниваем их
	if len(visibleWidgets) > 0 {
		y := (parentHeight - totalPreferredHeight) / 2
		for _, w := range visibleWidgets {
			_, _, wW, wH := w.Bounds()
			x := (parentWidth - wW) / 2 // Центрируем по ширине родителя
			w.SetBounds(x, y, wW, wH)
			y += wH + l.spacing
		}
	}
}
