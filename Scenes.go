package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	image  *ebiten.Image
	pushed bool
}

func NewButton(image *ebiten.Image) *Button {
	return &Button{image: image, pushed: false}
}

func (b *Button) IsPushed() bool {
	return b.pushed
}

func (b *Button) SetPushed(p bool) {
	b.pushed = p
}

func (b *Button) Draw(screen *ebiten.Image) {
	// Draw button implementation
}

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

	return nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	DrawTextAtCenter(screen, "Bullet Quest 2D")

}

func (m *MenuScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 128
}
