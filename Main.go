package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawTextAtCenter(screen *ebiten.Image, text string) {
	bounds := screen.Bounds()
	x := (bounds.Dx() - len(text)*7) / 2 // Approximate character width
	y := bounds.Dy() / 2
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

func init() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct {
	manager *SceneManager
}

func NewGame() *Game {
	g := &Game{}
	g.manager = &SceneManager{}
	g.manager.GoTo(NewMenuScene(g.manager))
	return g
}

func (g *Game) Update() error {
	return g.manager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.manager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.manager.Layout(outsideWidth, outsideHeight)
}

var PlayButtonNormal *CustomButton
var PlayButtonNormalPushed *CustomButton

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Bullet Quest 2D")
	buttonImageNormal := menuNormal
	buttonImagePushed := menuPushed
	PlayButtonNormal = NewCustomButton(100, 100, 64, 64, 1.0, buttonImageNormal)
	PlayButtonNormalPushed = NewCustomButton(100, 100, 64, 64, 1.0, buttonImagePushed)

	// Load character sprites used by the PlayScene
	LoadGameCharacters()

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

var (
	atlas, _, err = ebitenutil.NewImageFromFile("assets/bluebuttons.png")
	menuNormal    = atlas.SubImage(MenuNormalRect).(*ebiten.Image)
	menuPushed    = atlas.SubImage(MenuPushedRect).(*ebiten.Image) // if needed
	btn           = NewCustomButton(10, 10, 16, 16, 2.0, menuNormal)
)

func (g *Game) loadAssets() {
	if err != nil {
		log.Fatal(err)
	}

}
