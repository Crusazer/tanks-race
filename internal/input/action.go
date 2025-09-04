package input

type GameAction int

const (
	ActionMoveUp GameAction = iota
	ActionMoveDown
	ActionMoveLeft
	ActionMoveRight
	ActionShoot
	ActionLeftClick
	ActionPause
)

type UIAction int

const (
	UIActionNone UIAction = iota

	// Навигация
	UIActionNavigateUp   // Стрелка вверх
	UIActionNavigateDown // Стрелка вниз, Tab
	UIActionNavigatePrev // Shift + Tab (пока не реализуем, но заложим)

	// Действия
	UIActionSelect    // Выбор (клик мыши)
	UIActionConfirm   // Подтверждение (Enter)
	UIActionBackspace // Удаление символа (Backspace)
	UIActionBack      // Назад (Escape)
)
