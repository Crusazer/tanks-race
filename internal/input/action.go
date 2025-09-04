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
	UIActionUp
	UIActionDown
	UIActionSelect
	UIActionBack
)
