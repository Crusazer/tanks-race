package scenes

import (
	"image/color"
	"os"

	"github.com/Crusazer/tanks-race/internal/ui/screens"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuScene struct {
	menu         *screens.MainMenu
	sceneManager *SceneManager
}

func NewMenuScene(sceneManager *SceneManager) Scene {
	ms := &MenuScene{sceneManager: sceneManager}
	ms.menu = screens.NewMainMenu(ms.HandleManuAction)
	return ms
}

func (m *MenuScene) Enter() {}
func (m *MenuScene) Exit()  {}

func (m *MenuScene) Update(_ float64) error {
	return m.menu.Update()
}

func (m *MenuScene) Draw(dst *ebiten.Image) {
	dst.Fill(color.RGBA{40, 40, 60, 255})
	m.menu.Draw(dst)
}

func (m *MenuScene) Layout(w, h int) (int, int) { return w, h }

func (m *MenuScene) HandleManuAction(action screens.MenuAction) {
	switch action {
	case screens.ActionSinglePlay:
		m.sceneManager.ChangeScene(NewPlayingScene())
	case screens.ActionExit:
		os.Exit(0)
	}
}
