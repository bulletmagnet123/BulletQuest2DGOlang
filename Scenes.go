package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
	Enter()
	Exit()
}

type SceneManager struct {
	current Scene
}

func (sm *SceneManager) GoTo(s Scene) {
	if sm.current != nil {
		sm.current.Exit()
	}
	sm.current = s
	if sm.current != nil {
		sm.current.Enter()
	}
}

func (sm *SceneManager) Update() error {
	if sm.current == nil {
		return nil
	}
	return sm.current.Update()
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.current == nil {
		return
	}
	sm.current.Draw(screen)
}

func (sm *SceneManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	if sm.current == nil {
		return 320, 128
	}
	return sm.current.Layout(outsideWidth, outsideHeight)
}

// MenuScene: shows title and a Play button
type MenuScene struct {
	sm         *SceneManager
	wasPressed bool
}

func NewMenuScene(sm *SceneManager) *MenuScene {
	return &MenuScene{sm: sm}
}

func (m *MenuScene) Enter() {}
func (m *MenuScene) Exit()  {}

func (m *MenuScene) Update() error {
	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	mx, my := ebiten.CursorPosition()

	if StartGameButton != nil && pressed && StartGameButton.Contains(mx, my) {
		StartGameButton.SetPushed(true)
	} else {
		if StartGameButton != nil && StartGameButton.IsPushed() && m.wasPressed && StartGameButton.Contains(mx, my) {
			m.sm.GoTo(NewPlayScene(m.sm))
		}
		if StartGameButton != nil {
			StartGameButton.SetPushed(false)
		}
	}

	// Exit button: exit program when clicked
	if ExitGameButton != nil && pressed && ExitGameButton.Contains(mx, my) {
		ExitGameButton.SetPushed(true)
	} else {
		if ExitGameButton != nil && ExitGameButton.IsPushed() && m.wasPressed && ExitGameButton.Contains(mx, my) {
			os.Exit(0)
		}
		if ExitGameButton != nil {
			ExitGameButton.SetPushed(false)
		}
	}

	m.wasPressed = pressed
	return nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	DrawTextAtCenter(screen, "Bullet Quest 2D")

	if StartGameButton != nil {
		if StartGameButton.IsPushed() {
			StartGameButtonPushed.Draw(screen)
		} else {
			StartGameButton.Draw(screen)
		}
	}

	if ExitGameButton != nil {
		if ExitGameButton.IsPushed() {
			ExitGameButtonPushed.Draw(screen)
		} else {
			ExitGameButton.Draw(screen)
		}
	}
}

func (m *MenuScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 128
}
