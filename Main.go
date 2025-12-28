package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var PlayButtonNormal, PlayButtonNormalPushed, MenuNormalBtn, MenuPushedBtn *ebiten.Image

func DrawTextAtCenter(screen *ebiten.Image, text string) {
	bounds := screen.Bounds()
	x := (bounds.Dx() - len(text)*7) / 2 // Approximate character width
	y := bounds.Dy() / 2
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

func init() {
	// Load button atlas and initialize button images and button objects
	atlas, _, err := ebitenutil.NewImageFromFile("assets/bluebuttons.png")
	if err != nil {
		log.Fatal(err)
	}

	PlayButtonNormal = atlas.SubImage(PlayingNormalRect).(*ebiten.Image)
	PlayButtonNormalPushed = atlas.SubImage(PlayingPushedRect).(*ebiten.Image)
	MenuNormalBtn = atlas.SubImage(MenuNormalRect).(*ebiten.Image)
	MenuPushedBtn = atlas.SubImage(MenuPushedRect).(*ebiten.Image)

	// Place the Start and Exit buttons in the top-left of the screen
	StartGameButton = NewCustomButton(10, 10, 64, 32, 2.0, PlayButtonNormal)
	ExitGameButton = NewCustomButton(10, 60, 64, 32, 2.0, MenuNormalBtn)
	StartGameButtonPushed = NewCustomButton(10, 10, 64, 32, 2.0, PlayButtonNormalPushed)
	ExitGameButtonPushed = NewCustomButton(10, 60, 64, 32, 2.0, MenuPushedBtn)

}

type Game struct {
	manager *SceneManager
	Player  *Player
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
	screen.Fill(color.RGBA{120, 180, 255, 255}) // gray background
	if StartGameButton != nil {
		StartGameButton.Draw(screen)
	}
	g.manager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.manager.Layout(outsideWidth, outsideHeight)
}

var StartGameButton *CustomButton
var ExitGameButton *CustomButton
var StartGameButtonPushed *CustomButton
var ExitGameButtonPushed *CustomButton

func main() {
	ebiten.SetWindowTitle("Bullet Quest 2D")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(60)
	// Load character sprites used by the PlayScene
	LoadGameCharacters()
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}

}
