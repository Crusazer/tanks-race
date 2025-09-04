package screens

import (
	"fmt"
	"image/color"

	"github.com/Crusazer/tanks-race/internal/ui/core"
	"github.com/Crusazer/tanks-race/pkg/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuAction string

const (
	ActionSinglePlay  MenuAction = "single_play"
	ActionNetworkPlay MenuAction = "network_play"
	ActionEditor      MenuAction = "editor"
	ActionTankConfig  MenuAction = "tank_config"
	ActionSettings    MenuAction = "settings"
	ActionExit        MenuAction = "exit"
)

type MainMenu struct {
	layout   *core.VerticalFlowLayout
	onAction func(MenuAction)
}

func NewMainMenu(onAction func(MenuAction)) *MainMenu {
	layout := core.NewVerticalFlowLayout(30, 15)

	buttons := []string{
		"Игра по сети",
		"Редактор уровней",
		"Настройка танка",
		"Настройки",
	}

	singlePlayButton := createButton("Одиночная игра", func() { onAction(ActionSinglePlay) })
	layout.Add(singlePlayButton)

	for _, label := range buttons {
		btn := &core.Button{
			Text:         label,
			Width:        240,
			Height:       40,
			NormalColor:  color.RGBA{100, 100, 100, 255},
			HoverColor:   color.RGBA{150, 150, 150, 255},
			PressedColor: color.RGBA{200, 200, 200, 255},
			OnClick: func() {
				fmt.Println("Нажата:", label)
				// TODO: переключить сцену
			},
		}
		layout.Add(btn)
	}

	exitButton := createButton("Выход", func() { onAction(ActionExit) })
	layout.Add(exitButton)
	return &MainMenu{layout: layout, onAction: onAction}
}

func (m *MainMenu) Update() error {
	mx, my := ebiten.CursorPosition()
	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	mp := math.Vector2{X: float64(mx), Y: float64(my)}

	for _, w := range m.layout.Widgets() {
		w.Update(mp, pressed)
	}
	return nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	m.layout.ComputeBounds(w, h)

	for _, w := range m.layout.Widgets() {
		w.Draw(screen)
	}
}

func createButton(text string, callback func()) *core.Button {
	return &core.Button{
		Text:         text,
		Width:        240,
		Height:       40,
		NormalColor:  color.RGBA{100, 100, 100, 255},
		HoverColor:   color.RGBA{150, 150, 150, 255},
		PressedColor: color.RGBA{200, 200, 200, 255},
		OnClick:      callback,
	}
}
