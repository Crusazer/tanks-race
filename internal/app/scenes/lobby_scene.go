package scenes

import (
	"image/color"

	"github.com/Crusazer/tanks-race/internal/ui/core"
	"github.com/Crusazer/tanks-race/internal/ui/screens"
	"github.com/hajimehoshi/ebiten/v2"
)

type LobbyScene struct {
	lobby        *screens.Lobby
	sceneManager *SceneManager
}

func NewLobbyScene(sceneManager *SceneManager) Scene {
	ls := &LobbyScene{sceneManager: sceneManager}
	ls.lobby = screens.NewLobby(core.DefaultTheme)
	return ls
}

func (l *LobbyScene) Enter() {}
func (l *LobbyScene) Exit()  {}

func (l *LobbyScene) Update(_ float64) error {
	return l.lobby.Update()
}

func (l *LobbyScene) Draw(dst *ebiten.Image) {
	dst.Fill(color.RGBA{40, 40, 60, 255})
	l.lobby.Draw(dst)
}

func (l *LobbyScene) Layout(w, h int) (int, int) { return w, h }
